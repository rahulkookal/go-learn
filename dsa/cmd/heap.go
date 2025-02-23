/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"container/heap"
	"fmt"

	"github.com/spf13/cobra"
	mheap "rahul.com/dsa/heap"
)

// heapCmd represents the heap command
var heapCmd = &cobra.Command{
	Use:   "heap",
	Short: "Min Heap",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("heap called")
		h := &mheap.MinHeap{}
		heap.Init(h)
		arr := []int{3, 2, 1, 5, 6, 4}
		for _, i := range arr {
			heap.Push(h, i)
		}
		fmt.Println(heap.Pop(h))
	},
}

func init() {
	rootCmd.AddCommand(heapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// heapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// heapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
