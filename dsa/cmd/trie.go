/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/dsa/trie"
)

// trieCmd represents the trie command
var trieCmd = &cobra.Command{
	Use:   "trie",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trie called")
		t := trie.NewTrie()
		t.Insert("Rahul")
		t.Insert("Ram")
		t.Insert("Raman")
		t.Insert("Rahu")
		t.Traverse()
		fmt.Print(t.TraversePre("Rah"))
		if t.Search("Rahuls") {
			fmt.Println("Found")
		} else {
			fmt.Println("NOt Found")
		}
		fmt.Println("trie called")

	},
}

func init() {
	rootCmd.AddCommand(trieCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trieCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trieCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
