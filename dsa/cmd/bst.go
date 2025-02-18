/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/dsa/bst"
)

// dsaCmd represents the dsa command
var bstCmd = &cobra.Command{
	Use:   "bst",
	Short: "Binary Search Tree",
	Long:  `BST Implementation with Insert, PreOrderTraversal and InOrderTraversal`,
	Run: func(cmd *cobra.Command, args []string) {
		root := &bst.Node{Value: 50}
		// var a bst.Node = bst.Node{}
		root.Insert(10)
		root.Insert(20)
		root.Insert(5)
		root.Insert(2)
		root.InOrderTraversal()
		fmt.Print("\n")
		root.PreOrder()
		fmt.Print("\n")
	},
}

func init() {
	rootCmd.AddCommand(bstCmd)
}
