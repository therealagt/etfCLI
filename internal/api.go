package internal

import (
	"fmt"
	"net/http"
)

// Main Response Structure
type ETFData struct {
    MetaData   MetaData            `json:"Meta Data"`
    TimeSeries map[string]DayData  `json:"Time Series (Daily)"`
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

func fetchETFData(symbol string) (*ETFData, error) {
	resp, err := http.Get("https://alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=" + symbol + "&apikey=YOUR_API)
}