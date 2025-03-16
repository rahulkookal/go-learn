/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// experimentCmd represents the experiment command
var experimentCmd = &cobra.Command{
	Use:   "experiment",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("experiment called")
		var msg chan string = make(chan string, 2)
		go sendData(msg)

		var message string = <-msg
		var message1 string = <-msg
		fmt.Println("Recieved message: ", message)
		fmt.Println("Recieved message: ", message1)
	},
}

func sendData(ch chan string) {
	ch <- "Hello from Go Routines"
	ch <- "Hello from Go Routines1"
}
func init() {
	rootCmd.AddCommand(experimentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// experimentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// experimentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
