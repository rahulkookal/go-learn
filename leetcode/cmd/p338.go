/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// p338Cmd represents the p338 command
var p338Cmd = &cobra.Command{
	Use:   "p338",
	Short: "Counting Bits",
	Long: `Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

	

	Example 1:

	Input: n = 2
	Output: [0,1,1]
	Explanation:
	0 --> 0
	1 --> 1
	2 --> 10
	Example 2:

	Input: n = 5
	Output: [0,1,1,2,1,2]
	Explanation:
	0 --> 0
	1 --> 1
	2 --> 10
	3 --> 11
	4 --> 100
	5 --> 101
	

	Constraints:

	0 <= n <= 105
	

	Follow up:

	It is very easy to come up with a solution with a runtime of O(n log n). Can you do it in linear time O(n) and possibly in a single pass?
	Can you do it without using any built-in function (i.e., like __builtin_popcount in C++)?`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("p338 called")
		result := solutions.CountBits(5)
		fmt.Println("Bit counts for 5: ", result)
	},
}

func init() {
	rootCmd.AddCommand(p338Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// p338Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// p338Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
