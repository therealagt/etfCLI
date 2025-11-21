package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therealagt/etfcli/internal"
)

var infoCmd = &cobra.Command{
	Use:   "info [ETF_SYMBOL]",
	Short: "Analyze a specific ETF",
	Long:  `Analyze and provide detailed information about a specific ETF identified by its symbol.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide one ETF symbol.")
			return
		}

		data, err := internal.FetchETFDataCached(args[0])
		if err != nil {
			fmt.Printf("Error fetching data for %s: %v\n", args[0], err)
			return
		}

		fmt.Printf("\nETF Information for %s\n", data.Symbol)
		fmt.Println("==========================================")
		fmt.Printf("Symbol: %s\n", data.Symbol)
		fmt.Printf("Current Price: $%.2f\n", data.Price)
		fmt.Printf("Previous Close: $%.2f\n", data.PreviousClose)
		fmt.Printf("Change: $%.2f (%.2f%%)\n", data.Change, data.ChangePercent)
		fmt.Printf("Volume: %d\n", data.Volume)
		fmt.Printf("Last Updated: %s\n", data.LastUpdated.Format("2006-01-02 15:04:05"))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
