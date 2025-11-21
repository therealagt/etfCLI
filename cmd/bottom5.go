package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therealagt/etfcli/internal"
)

var bottom5Cmd = &cobra.Command{
	Use:   "bottom5",
	Short: "List the bottom five performing ETFs",
	Long:  `Display the bottom five performing ETFs based on daily percentage change.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Loading ETF symbols...")

		symbols, err := internal.LoadETFSymbols()
		if err != nil {
			fmt.Printf("Error loading ETF symbols: %v\n", err)
			return
		}

		fmt.Printf("Analyzing %d ETFs for worst performers...\n", len(symbols))

		etfPerformances, err := internal.GetBottomETFsWithPerformance(symbols)
		if err != nil {
			fmt.Printf("Error getting ETF performance data: %v\n", err)
			return
		}

		if len(etfPerformances) == 0 {
			fmt.Println("No ETF performance data available")
			return
		}

		bottomCount := 5
		if len(etfPerformances) < bottomCount {
			bottomCount = len(etfPerformances)
		}

		fmt.Printf("\nBottom %d Performing ETFs:\n", bottomCount)
		fmt.Println("==========================================")

		for i := 0; i < bottomCount; i++ {
			etf := etfPerformances[i]
			fmt.Printf("%d. %s - %.2f%%\n", i+1, etf.Symbol, etf.Performance)
			fmt.Printf("   Current Price: $%.2f\n", etf.Data.Price)
			fmt.Printf("   Change: $%.2f\n", etf.Data.Change)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(bottom5Cmd)
}
