package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therealagt/etfcli/internal"
)

var top5Cmd = &cobra.Command{
	Use:   "top5",
	Short: "List the top five performing ETFs",
	Long:  `Display the top five performing ETFs based on daily percentage change.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Loading ETF symbols...")

		symbols, err := internal.LoadETFSymbols()
		if err != nil {
			fmt.Printf("Error loading ETF symbols: %v\n", err)
			return
		}

		fmt.Printf("Analyzing %d ETFs for top performers...\n", len(symbols))

		etfPerformances, err := internal.GetTopETFsWithPerformance(symbols)
		if err != nil {
			fmt.Printf("Error getting ETF performance data: %v\n", err)
			return
		}

		if len(etfPerformances) == 0 {
			fmt.Println("No ETF performance data available")
			return
		}

		topCount := 5
		if len(etfPerformances) < topCount {
			topCount = len(etfPerformances)
		}

		fmt.Printf("\nTop %d Performing ETFs:\n", topCount)
		fmt.Println("==========================================")

		for i := 0; i < topCount; i++ {
			etf := etfPerformances[i]
			fmt.Printf("%d. %s - %.2f%%\n", i+1, etf.Symbol, etf.Performance)
			fmt.Printf("   Current Price: $%.2f\n", etf.Data.Price)
			fmt.Printf("   Change: $%.2f\n", etf.Data.Change)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(top5Cmd)
}
