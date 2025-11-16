package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "etfcli",
		Short: "ETF CLI is a command-line tool for analyzing ETFs",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to ETF CLI! Use --help to see available commands.")

			fmt.Println("Top five ETFs:")
			//logic
			fmt.Println("Bottom five ETFs:")
			//logic
			fmt.Println("Use 'etfcli info [ETF_SYMBOL]' to analyze a specific ETF.")
			//logic
		},
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "top5",
		Short: "List the top five performing ETFs",
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "bottom5",
		Short: "List the bottom five performing ETFs",
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "info [ETF_SYMBOL]",
		Short: "Analyze a specific ETF",
	})

	rootCmd.Execute()
}
