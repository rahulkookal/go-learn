/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rahul.com/leetcode/solutions"
)

// 435Cmd represents the 435 command
var p435Cmd = &cobra.Command{
	Use:   "435",
	Short: "Non-overlapping Intervals",
	Long: `Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

		Note that intervals which only touch at a point are non-overlapping. For example, [1, 2] and [2, 3] are non-overlapping.

		

		Example 1:

		Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
		Output: 1
		Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
		Example 2:

		Input: intervals = [[1,2],[1,2],[1,2]]
		Output: 2
		Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.
		Example 3:

		Input: intervals = [[1,2],[2,3]]
		Output: 0
		Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
		

		Constraints:

		1 <= intervals.length <= 105
		intervals[i].length == 2
		-5 * 104 <= starti < endi <= 5 * 104`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("435 called")
		result := solutions.EraseOverlapIntervals([][]int{{1, 2}, {2, 3}})
		fmt.Println("No overlapping intervals: ", result)
	},
}

func init() {
	rootCmd.AddCommand(p435Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 435Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 435Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
