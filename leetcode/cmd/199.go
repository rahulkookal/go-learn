/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// 199Cmd represents the 199 command
var p199Cmd = &cobra.Command{
	Use:   "199",
	Short: "Binary Tree Right Side View",
	Long: `Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
	Example 1:
	Input: root = [1,2,3,null,5,null,4]
	Output: [1,3,4]
	Explanation:
	Example 2:
	Input: root = [1,2,3,4,null,null,null,5]
	Output: [1,3,4,5]
	Explanation:

	Example 3:
	Input: root = [1,null,3]
	Output: [1,3]
	Example 4:
	Input: root = []
	Output: []
	Constraints:
	The number of nodes in the tree is in the range [0, 100].
	-100 <= Node.val <= 100`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("199 called")
		result := solutions.RightSideView(&solutions.TreeNode{
			Val:  10,
			Left: nil,
			Right: &solutions.TreeNode{
				Val:   12,
				Left:  nil,
				Right: nil},
		})
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(p199Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 199Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 199Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
