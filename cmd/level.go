/*
Copyright © 2023 hawk87
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// levelCmd represents the params command
var levelCmd = &cobra.Command{
	Use:   "level",
	Short: "Set the level and get EXP and collection number",
	Long: "Set the level and get EXP and collection number\n" +
		"PLEASE NOTE: The collection field is subtracted by 1.",
	Run: func(cmd *cobra.Command, args []string) {
		level, err := cmd.Flags().GetInt("number")
		if err != nil {
			return
		}
		if level == 0 {
			fmt.Println("Please specify a level number. e.g.: state-fixer level -n 10")
			return
		}
		fmt.Printf("exp: %d\n", calculateExp(level))
		fmt.Printf("i: %d\n", calculateI(level)-1)
	},
}

func init() {
	rootCmd.AddCommand(levelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// levelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// levelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	levelCmd.Flags().IntP("number", "n", 0, "Set level number")
}

var delta2 = 90

func calculateDelta1(n int) int {
	if n == 1 {
		return 1
	}
	return calculateDelta1(n-1) + delta2
}

func calculateExp(n int) int {
	if n == 1 {
		return 1
	}
	return calculateExp(n-1) + calculateDelta1(n)
}

func calculateI(n int) int {
	if n == 1 {
		return 1
	}
	return calculateI(n-1) + 3
}
