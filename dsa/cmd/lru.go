/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/dsa/lru"
)

// lruCmd represents the lru command
var lruCmd = &cobra.Command{
	Use:   "lru",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lru called")
		lru := lru.Constructor(2) // Cache with capacity = 2

		lru.Put(1, 100)
		lru.Put(2, 200)
		fmt.Println(lru.Get(1)) // Returns 100 (key 1 is now most recently used)

		lru.Put(3, 300)         // Evicts key 2 (LRU is removed)
		fmt.Println(lru.Get(2)) // Returns -1 (not found)

		lru.Put(4, 400)         // Evicts key 1
		fmt.Println(lru.Get(1)) // Returns -1
		fmt.Println(lru.Get(3)) // Returns 300
		fmt.Println(lru.Get(4))
	},
}

func init() {
	rootCmd.AddCommand(lruCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lruCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lruCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
