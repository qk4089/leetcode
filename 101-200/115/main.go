package _115

//https://leetcode.cn/problems/distinct-subsequences/
//给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数，结果需要对 10^9 + 7 取模。

//示例 1：
//输入：s = "rabbbit", t = "rabbit"
//输出：3
//解释：
//如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
//rabbbit
//rabbbit
//rabbbit

//输入：s = "babgbag", t = "bag"
//输出：5
//解释：
//如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
//babgbag
//babgbag
//babgbag
//babgbag
//babgbag

//提示：
//	1 <= s.length, t.length <= 1000
//	s 和 t 由英文字母组成

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= i && j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
