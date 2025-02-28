/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// 1207Cmd represents the 1207 command
var p1207Cmd = &cobra.Command{
	Use:   "1207",
	Short: "Unique Number of Occurrences",
	Long: `Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.
		Example 1:

		Input: arr = [1,2,2,1,1,3]
		Output: true
		Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
		Example 2:

		Input: arr = [1,2]
		Output: false
		Example 3:

		Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
		Output: true
 `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1207 called")
		result := solutions.UniqueOccurrences([]int{1, 2, 3, 2, 3, 3})
		fmt.Println("Result: ", result)
	},
}

func init() {
	rootCmd.AddCommand(p1207Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 1207Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 1207Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
