/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// 724Cmd represents the 724 command
var p724Cmd = &cobra.Command{
	Use:   "724",
	Short: "Find Pivot Index",
	Long: `Given an array of integers nums, calculate the pivot index of this array.

	The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

	If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

	Return the leftmost pivot index. If no such index exists, return -1.

	

	Example 1:

	Input: nums = [1,7,3,6,5,6]
	Output: 3
	Explanation:
	The pivot index is 3.
	Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
	Right sum = nums[4] + nums[5] = 5 + 6 = 11
	Example 2:

	Input: nums = [1,2,3]
	Output: -1
	Explanation:
	There is no index that satisfies the conditions in the problem statement.
	Example 3:

	Input: nums = [2,1,-1]
	Output: 0
	Explanation:
	The pivot index is 0.
	Left sum = 0 (no elements to the left of index 0)
	Right sum = nums[1] + nums[2] = 1 + -1 = 0`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("724 called")
		in := []int{1, 7, 3, 6, 5, 6}
		result := solutions.PivotIndex(in)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(p724Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 724Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 724Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
