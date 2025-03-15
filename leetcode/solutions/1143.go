package solutions

import "fmt"

func LongestCommonSubsequence(text1, text2 string) int {
	if text1 == text2 {
		return len(text1)
	}
	m, n := len(text1), len(text2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Print(dp)
	return dp[m][n]
}
