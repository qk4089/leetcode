package _1312

//https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/
//给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
//请你返回让 s 成为回文串的 最少操作次数 。
//「回文串」是正读和反读都相同的字符串。

//示例 1：
//输入：s = "zzazz"
//输出：0
//解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。

//输入：s = "mbadm"
//输出：2
//解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。

//输入：s = "leetcode"
//输出：5
//解释：插入 5 个字符后字符串变为 "leetcodocteel" 。

//提示：
//	1 <= s.length <= 500
//	s 中所有字符都是小写字母。

func minInsertions(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return len(s) - dp[0][len(s)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
