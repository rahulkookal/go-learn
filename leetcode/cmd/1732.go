/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// 1732Cmd represents the 1732 command
var p1732Cmd = &cobra.Command{
	Use:   "1732",
	Short: "Find the Highest Altitude",
	Long: `There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes. 
	The biker starts his trip on point 0 with altitude equal 0.

	You are given an integer array gain of length n where gain[i] is the net gain in altitude between points 
	and i + 1 for all (0 <= i < n). Return the highest altitude of a point.

	

	Example 1:

	Input: gain = [-5,1,5,0,-7]
	Output: 1
	Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
	Example 2:

	Input: gain = [-4,-3,-2,-1,4,3,2]
	Output: 0
	Explanation: The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.
	

	Constraints:

	n == gain.length
	1 <= n <= 100
	-100 <= gain[i] <= 100`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1732 called")
		result := solutions.LargestAltitude([]int{-4, -3, -2, -1, 4, 3, 2})
		fmt.Print("Result: ", result)
	},
}

func init() {
	rootCmd.AddCommand(p1732Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 1732Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 1732Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
