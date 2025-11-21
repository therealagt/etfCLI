package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// Alpha Vantage API Response Structure
type AlphaVantageResponse struct {
	MetaData   MetaData                       `json:"Meta Data"`
	TimeSeries map[string]AlphaVantageDayData `json:"Time Series (Daily)"`
}

type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
}

type AlphaVantageDayData struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

// Our simplified ETF Data Structure
type ETFData struct {
	Symbol        string    `json:"symbol"`
	Name          string    `json:"name"`
	Price         float64   `json:"price"`
	PreviousClose float64   `json:"previous_close"`
	Change        float64   `json:"change"`
	ChangePercent float64   `json:"change_percent"`
	Volume        int64     `json:"volume"`
	LastRefreshed string    `json:"last_refreshed"`
	LastUpdated   time.Time `json:"last_updated"`
}

// Cache structure
type CacheData struct {
	Data      *ETFData  `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func init() {
	godotenv.Load()
	os.MkdirAll("cache", 0755)
}

// Load ETF symbols
func LoadETFSymbols() ([]string, error) {
	data, err := os.ReadFile("etf_symbols.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read etf_symbols.json: %v", err)
	}

	var symbols []string
	err = json.Unmarshal(data, &symbols)
	if err != nil {
		return nil, fmt.Errorf("failed to parse etf_symbols.json: %v", err)
	}

	return symbols, nil
}

// Cache functions
func getCacheFilePath(symbol string) string {
	today := time.Now().Format("2006-01-02")
	return filepath.Join("cache", fmt.Sprintf("%s_%s.json", symbol, today))
}

func loadFromCache(symbol string) *ETFData {
	cachePath := getCacheFilePath(symbol)
	data, err := os.ReadFile(cachePath)
	if err != nil {
		return nil
	}

	var cached CacheData
	if err := json.Unmarshal(data, &cached); err != nil {
		return nil
	}

	// Check if cache is still valid (same day)
	if cached.Timestamp.Format("2006-01-02") == time.Now().Format("2006-01-02") {
		return cached.Data
	}

	return nil
}

func saveToCache(symbol string, data *ETFData) {
	cachePath := getCacheFilePath(symbol)

	cached := CacheData{
		Data:      data,
		Timestamp: time.Now(),
	}

	jsonData, err := json.MarshalIndent(cached, "", "  ")
	if err != nil {
		return
	}

	os.WriteFile(cachePath, jsonData, 0644)
}

// Fetch ETF data from Alpha Vantage
func FetchETFData(symbol string) (*ETFData, error) {
	apiKey := os.Getenv("ALPHAVANTAGE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ALPHAVANTAGE_API_KEY not found in environment")
	}

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", symbol, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data for %s: %v", symbol, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d for symbol %s", resp.StatusCode, symbol)
	}

	var avResponse AlphaVantageResponse
	if err := json.NewDecoder(resp.Body).Decode(&avResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response for %s: %v", symbol, err)
	}

	if len(avResponse.TimeSeries) == 0 {
		return nil, fmt.Errorf("no data found for symbol %s", symbol)
	}

	// Get the most recent day's data
	var latestDate string
	for date := range avResponse.TimeSeries {
		if latestDate == "" || date > latestDate {
			latestDate = date
		}
	}

	latestData := avResponse.TimeSeries[latestDate]

	// Parse prices
	var price, prevClose float64
	var volume int64
	fmt.Sscanf(latestData.Close, "%f", &price)

	// Get previous day for change calculation
	var previousDate string
	for date := range avResponse.TimeSeries {
		if date != latestDate && (previousDate == "" || date > previousDate) {
			previousDate = date
		}
	}

	if previousDate != "" {
		fmt.Sscanf(avResponse.TimeSeries[previousDate].Close, "%f", &prevClose)
	}

	fmt.Sscanf(latestData.Volume, "%d", &volume)

	// Calculate change
	change := price - prevClose
	changePercent := 0.0
	if prevClose > 0 {
		changePercent = (change / prevClose) * 100
	}

	// Convert to our ETFData structure
	data := &ETFData{
		Symbol:        avResponse.MetaData.Symbol,
		Name:          symbol, // Alpha Vantage doesn't provide name in this endpoint
		Price:         price,
		PreviousClose: prevClose,
		Change:        change,
		ChangePercent: changePercent,
		Volume:        volume,
		LastRefreshed: avResponse.MetaData.LastRefreshed,
		LastUpdated:   time.Now(),
	}

	return data, nil
}

// Fetch with caching
func FetchETFDataCached(symbol string) (*ETFData, error) {
	if cachedData := loadFromCache(symbol); cachedData != nil {
		fmt.Printf("Using cached data for %s\n", symbol)
		return cachedData, nil
	}

	fmt.Printf("Fetching fresh data for %s from Alpha Vantage\n", symbol)
	data, err := FetchETFData(symbol)
	if err != nil {
		return nil, err
	}

	saveToCache(symbol, data)
	return data, nil
}

// Fetch multiple ETFs with caching
func FetchMultipleETFsCached(symbols []string) ([]*ETFData, error) {
	fmt.Printf("Fetching data for %d ETFs (with caching)...\n", len(symbols))
	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make([]*ETFData, 0, len(symbols))
	errors := make([]error, 0)

	for _, symbol := range symbols {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()

			data, err := FetchETFDataCached(s)

			mu.Lock()
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to fetch %s: %v", s, err))
			} else {
				results = append(results, data)
			}
			mu.Unlock()
		}(symbol)
	}

	wg.Wait()

	if len(errors) > 0 {
		fmt.Printf("Some requests failed: %v\n", errors)
		if len(results) == 0 {
			return nil, fmt.Errorf("all requests failed")
		}
	}

	fmt.Printf("Successfully fetched %d out of %d ETFs\n", len(results), len(symbols))
	return results, nil
}
