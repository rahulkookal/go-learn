/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// problem104Cmd represents the problem104 command
var problem104Cmd = &cobra.Command{
	Use:   "104",
	Short: "Maximum Depth of Binary Tree",
	Long: `Given the root of a binary tree, return its maximum depth.

	A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

	

	Example 1:


	Input: root = [3,9,20,null,null,15,7]
	Output: 3
	Example 2:

	Input: root = [1,null,2]
	Output: 2
	

	Constraints:

	The number of nodes in the tree is in the range [0, 104].
	-100 <= Node.val <= 100`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("problem104 called")
		root := &solutions.TreeNode{Val: 1}
		depth := solutions.MaxDepth(root)
		fmt.Println("Max depth: ", depth)
	},
}

func init() {
	rootCmd.AddCommand(problem104Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// problem104Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// problem104Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
