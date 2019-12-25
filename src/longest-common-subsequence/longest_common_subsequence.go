package longest_common_subsequence
// Bottom Up
func LongestCommonSubsequence(text1 string, text2 string) int {
	tlen := len(text1) + 1
	t2len := len(text2) + 1

	dp := make([][]int, tlen)

	for i, _ := range dp {
		dp[i] = make([]int, t2len)
	}

	for i:=1; i < tlen; i ++ {
		for j:=1; j< t2len; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}else{
				dp[i][j] = dp[i][j-1]
				if dp[i-1][j] > dp[i][j] {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[tlen-1][t2len-1]
}

// 思路：
// 1. 分析：假设 text1 和 text2 最后一个字符相同， 那么公共字符串 = text1[i-1] + text2[i-1] + 1;
//          若 最后一个字符串不相等公共字符串 = text1[i-1] + text2[i-1]
//
