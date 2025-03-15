/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// 1143Cmd represents the 1143 command
var p1143Cmd = &cobra.Command{
	Use:   "1143",
	Short: "Longest Common Subsequence",
	Long: `Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

	A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
	For example, "ace" is a subsequence of "abcde".
	A common subsequence of two strings is a subsequence that is common to both strings.
	Example 1:

	Input: text1 = "abcde", text2 = "ace" 
	Output: 3  
	Explanation: The longest common subsequence is "ace" and its length is 3.
	Example 2:

	Input: text1 = "abc", text2 = "abc"
	Output: 3
	Explanation: The longest common subsequence is "abc" and its length is 3.
	Example 3:

	Input: text1 = "abc", text2 = "def"
	Output: 0
	Explanation: There is no such common subsequence, so the result is 0.
	

	Constraints:

	1 <= text1.length, text2.length <= 1000
	text1 and text2 consist of only lowercase English characters.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1143 called")
		result := solutions.LongestCommonSubsequence("abc", "def")
		fmt.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(p1143Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 1143Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 1143Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
