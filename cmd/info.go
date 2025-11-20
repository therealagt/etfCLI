/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therealagt/etfcli/internal"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Analyze a specific ETF",
	Long:  `Analyze and provide detailed information about a specific ETF identified by its symbol.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide at least one ETF symbol.")
			return
		}

		data, err := internal.FetchETFData(args[0])
		if err != nil {
			fmt.Printf("Error fetching data for %s: %v\n", args[0], err)
			return
		}

		fmt.Printf("ETF Symbol: %s\n", data.MetaData.Symbol)
		fmt.Printf("Last Refreshed: %s\n", data.MetaData.LastRefreshed)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
