/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 1161Cmd represents the 1161 command
var p1161Cmd = &cobra.Command{
	Use:   "1161",
	Short: "Maximum Level Sum of a Binary Tree",
	Long: `Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

	Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

	Example 1:

	Input: root = [1,7,0,7,-8,null,null]
	Output: 2
	Explanation: 
	Level 1 sum = 1.
	Level 2 sum = 7 + 0 = 7.
	Level 3 sum = 7 + -8 = -1.
	So we return the level with the maximum sum which is level 2.
	Example 2:

	Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
	Output: 2
	
	Constraints:

	The number of nodes in the tree is in the range [1, 104].
	-105 <= Node.val <= 105`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1161 called")
	},
}

func init() {
	rootCmd.AddCommand(p1161Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 1161Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 1161Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
