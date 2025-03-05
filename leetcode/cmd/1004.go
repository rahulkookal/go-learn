/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// 1004Cmd represents the 1004 command
var p1004Cmd = &cobra.Command{
	Use:   "1004",
	Short: "Max Consecutive Ones III",
	Long: `Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

	Example 1:

	Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
	Output: 6
	Explanation: [1,1,1,0,0,1,1,1,1,1,1]
	Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
	Example 2:

	Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
	Output: 10
	Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
	Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
	

	Constraints:

	1 <= nums.length <= 105
	nums[i] is either 0 or 1.
	0 <= k <= nums.length`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1004 called")
		in := []int{1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1}
		expected := 9
		result := solutions.LongestOnes(in, 2)
		fmt.Printf("LetterCombinations(%d) = %d\n expected %d \n", in, result, expected)
	},
}

func init() {
	rootCmd.AddCommand(p1004Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 1004Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 1004Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
