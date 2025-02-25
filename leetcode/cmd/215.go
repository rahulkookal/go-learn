/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// p215Cmd represents the p215 command
var p215Cmd = &cobra.Command{
	Use:   "215",
	Short: "Kth Largest Element in an Array",
	Long: `Given an integer array nums and an integer k, return the kth largest element in the array.
		Note that it is the kth largest element in the sorted order, not the kth distinct element.
		Can you solve it without sorting?
		Example 1:
		Input: nums = [3,2,1,5,6,4], k = 2
		Output: 5
		Example 2:
		Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
		Output: 4
		Constraints:
		1 <= k <= nums.length <= 105
		-104 <= nums[i] <= 104`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("p215 called")
		arr := []int{1, 2, 3, 4}
		fmt.Println("Kth Largest number: ", solutions.KThLargestNumber(arr, 2))
	},
}

func init() {
	rootCmd.AddCommand(p215Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// p215Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// p215Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
