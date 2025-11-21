# ETF CLI

A command-line tool for analyzing Exchange-Traded Fund (ETF) performance using real-time market data from Alpha Vantage API. The application provides three core commands: `info` for detailed information about a specific ETF, `top5` to display the best performing ETFs, and `bottom5` to show the worst performers based on daily percentage change. Built with Go and the Cobra CLI framework, the tool includes intelligent caching to minimize API calls and supports concurrent data fetching for efficient batch analysis of multiple ETF symbols.

## Usage

```bash
# Get detailed information about a specific ETF
./etfcli info SPY

# Show top 5 performing ETFs
./etfcli top5

# Show bottom 5 performing ETFs
./etfcli bottom5
```

## Setup

1. Get your free API key from Alpha Vantage
2. Create a `.env` file with `ALPHAVANTAGE_API_KEY=your_key_here`
3. Run `go build` to compile or `go run main.go [command]` to execute directly

## Requirements

- Go 1.21+
- Alpha Vantage API key (free tier: 500 calls/day)