/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// 17Cmd represents the 17 command
var p17Cmd = &cobra.Command{
	Use:   "17",
	Short: "Letter Combinations of a Phone Number",
	Long: `Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

	A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
	
	Example 1:

	Input: digits = "23"
	Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	Example 2:

	Input: digits = ""
	Output: []
	Example 3:

	Input: digits = "2"
	Output: ["a","b","c"]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("17 called")
		in := "23"
		result := solutions.LetterCombinations(in)
		expectedOut := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
		fmt.Printf("LetterCombinations(%q) = %q\n expected %q \n", in, result, expectedOut)
	},
}

func init() {
	rootCmd.AddCommand(p17Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 17Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 17Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
