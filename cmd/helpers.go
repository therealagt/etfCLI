package cmd

import "fmt"

// Helper function to format large numbers (shared across commands)
func formatNumber(num int64) string {
	if num == 0 {
		return "N/A"
	}

	if num >= 1000000000 {
		return fmt.Sprintf("%.2fB", float64(num)/1000000000)
	} else if num >= 1000000 {
		return fmt.Sprintf("%.2fM", float64(num)/1000000)
	} else if num >= 1000 {
		return fmt.Sprintf("%.2fK", float64(num)/1000)
	}
	return fmt.Sprintf("%d", num)
}

// Helper function to format volume numbers (shared across commands)
func formatVolume(num int64) string {
	if num == 0 {
		return "N/A"
	}

	if num >= 1000000000 {
		return fmt.Sprintf("%.1fB", float64(num)/1000000000)
	} else if num >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(num)/1000000)
	} else if num >= 1000 {
		return fmt.Sprintf("%.1fK", float64(num)/1000)
	}
	return fmt.Sprintf("%d", num)
}
