/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// p1456Cmd represents the p1456 command
var p1456Cmd = &cobra.Command{
	Use:   "p1456",
	Short: "Maximum Number of Vowels in a Substring of Given Length",
	Long: `Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

	Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

	

	Example 1:

	Input: s = "abciiidef", k = 3
	Output: 3
	Explanation: The substring "iii" contains 3 vowel letters.
	Example 2:

	Input: s = "aeiou", k = 2
	Output: 2
	Explanation: Any substring of length 2 contains 2 vowels.
	Example 3:

	Input: s = "leetcode", k = 3
	Output: 2
	Explanation: "lee", "eet" and "ode" contain 2 vowels.
	

	Constraints:

	1 <= s.length <= 105
	s consists of lowercase English letters.
	1 <= k <= s.length`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("p1456 called")
		var res int = solutions.MaxVowels("leetcode", 3)
		fmt.Println("Result: ", res)
	},
}

func init() {
	rootCmd.AddCommand(p1456Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// p1456Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// p1456Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
