package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// Main Response Structure
type ETFData struct {
	MetaData   MetaData           `json:"Meta Data"`
	TimeSeries map[string]DayData `json:"Time Series (Daily)"`
}

// Meta Information
type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
}

// Dayly Data Structure
type DayData struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func init() {
	godotenv.Load()
}

// Fetch ETF data from Alpha Vantage API
func FetchETFData(symbol string) (*ETFData, error) {
	apiKey := os.Getenv("ALPHAVANTAGE_API_KEY")
	fmt.Printf("API Key was loaded as: %s\n", apiKey)
	if apiKey == "" {
		return nil, fmt.Errorf("API key not set in environment variable ALPHAVANTAGE_API_KEY")
	}

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", symbol, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch data: %s", resp.Status)
	}

	var data ETFData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Load multiple ETFs concurrently
func FetchMultipleETFs(symbols []string) ([]*ETFData, error) {
	fmt.Printf("Fetching Data for %d ETFs...\n", len(symbols))
	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make([]*ETFData, 0, len(symbols))
	errors := make([]error, 0)

	for _, symbol := range symbols {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()

			data, err := FetchETFData(s)

			mu.Lock()
			if err != nil {
				errors = append(errors, fmt.Errorf("Failed to fetch %s: %v", s, err))
			} else {
				results = append(results, data)
			}
			mu.Unlock()
		}(symbol)
	}

	wg.Wait()

	if len(errors) > 0 {
		return results, fmt.Errorf("Some requests failed: %v", errors)
	}

	fmt.Printf("All %d requests successful\n", len(results))
	return results, nil
}
