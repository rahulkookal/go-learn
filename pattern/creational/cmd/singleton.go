/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/create/singleton"
)

// singletonCmd represents the singleton command
var singletonCmd = &cobra.Command{
	Use:   "singleton",
	Short: "A brief description of your command",
	Long:  `Sample implementation for singleton`,
	Run: func(cmd *cobra.Command, args []string) {
		var i *singleton.Singleton = singleton.GetSingletonInstance()
		var k *singleton.Singleton = singleton.GetSingletonInstance()
		if i == k {
			fmt.Println("singleton called")
			fmt.Println(k)
		} else {
			fmt.Println("singleton called2")
		}

	},
}

func init() {
	rootCmd.AddCommand(singletonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// singletonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// singletonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
