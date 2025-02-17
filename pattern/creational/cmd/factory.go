/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"rahul.com/create/factory"
)

// factoryCmd represents the factory command
var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "A brief description of your command",
	Long: `The Factory Method Pattern is a creational design pattern that provides an
	 	interface for creating objects but allows subclasses or 
	 	specific methods to define the actual instantiation.`,
	Run: func(cmd *cobra.Command, args []string) {
		var c factory.Shape = factory.GetShape(factory.CircleShape)
		c.Draw()
		var r factory.Shape = factory.GetShape(factory.RectangleShape)
		r.Draw()
	},
}

func init() {
	rootCmd.AddCommand(factoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// factoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// factoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
