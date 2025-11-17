/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// bottom5Cmd represents the bottom5 command
var bottom5Cmd = &cobra.Command{
	Use:   "bottom5",
	Short: "List the bottom five performing ETFs",
	Long: `The chosen 5 ETFS with the lowest performance over a specified period.
	This command fetches and displays the bottom five ETFs based on their performance metrics.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bottom5 called")
	},
}

func init() {
	rootCmd.AddCommand(bottom5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bottom5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bottom5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
