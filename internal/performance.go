package internal

import (
	"fmt"
	"sort"
)

// ETF with performance data
type ETFPerformance struct {
	Data        *ETFData
	Performance float64
	Symbol      string
}

// Calculate performance
func CalculatePerformance(data *ETFData) (float64, error) {
	return data.ChangePercent, nil
}

// Get ETFs with performance data (internal helper)
func GetETFsWithPerformance(symbols []string) ([]ETFPerformance, error) {
	etfDataList, err := FetchMultipleETFsCached(symbols)
	if err != nil {
		return nil, err
	}

	var etfPerformances []ETFPerformance
	for _, data := range etfDataList {
		performance, err := CalculatePerformance(data)
		if err != nil {
			fmt.Printf("Warning: Could not calculate performance for %s: %v\n", data.Symbol, err)
			continue
		}

		etfPerformances = append(etfPerformances, ETFPerformance{
			Data:        data,
			Performance: performance,
			Symbol:      data.Symbol,
		})
	}

	return etfPerformances, nil
}

// Get top ETFs
func GetTopETFsWithPerformance(symbols []string) ([]ETFPerformance, error) {
	etfPerformances, err := GetETFsWithPerformance(symbols)
	if err != nil {
		return nil, err
	}

	// Sort descending (best first)
	sort.Slice(etfPerformances, func(i, j int) bool {
		return etfPerformances[i].Performance > etfPerformances[j].Performance
	})

	return etfPerformances, nil
}

func GetBottomETFsWithPerformance(symbols []string) ([]ETFPerformance, error) {
	etfPerformances, err := GetETFsWithPerformance(symbols)
	if err != nil {
		return nil, err
	}

	// Sort ascending (worst first)
	sort.Slice(etfPerformances, func(i, j int) bool {
		return etfPerformances[i].Performance < etfPerformances[j].Performance
	})

	return etfPerformances, nil
}
