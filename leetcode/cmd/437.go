/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// p437Cmd represents the p437 command
var p437Cmd = &cobra.Command{
	Use:   "437",
	Short: "Path Sum III",
	Long: `Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

	The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

	

	Example 1:


	Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
	Output: 3
	Explanation: The paths that sum to 8 are shown.
	Example 2:

	Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
	Output: 3`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("p437 called")
		root := &solutions.TreeNode{Val: 1}
		out := solutions.PathSum(root, 1)
		fmt.Println("out: ", out)
	},
}

func init() {
	rootCmd.AddCommand(p437Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// p437Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// p437Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
