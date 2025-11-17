/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// top5Cmd represents the top5 command
var top5Cmd = &cobra.Command{
	Use:   "top5",
	Short: "List the top five performing ETFs",
	Long: `The chosen 5 ETFS with the highest performance over a specified period.
	This command fetches and displays the top five ETFs based on their performance metrics.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("top5 called")
	},
}

func init() {
	rootCmd.AddCommand(top5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// top5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// top5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
