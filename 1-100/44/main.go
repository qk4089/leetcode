package _44

//https://leetcode.cn/problems/wildcard-matching/
//给你一个输入字符串 (s) 和一个字符模式 (p) ，请你实现一个支持 '?' 和 '*' 匹配规则的通配符匹配：
//	'?' 可以匹配任何单个字符。
//	'*' 可以匹配任意字符序列（包括空字符序列）。
//判定匹配成功的充要条件是：字符模式必须能够 完全匹配 输入字符串（而不是部分匹配）。

//示例 1：
//输入：s = "aa", p = "a"
//输出：false
//解释："a" 无法匹配 "aa" 整个字符串。

//输入：s = "aa", p = "*"
//输出：true
//解释：'*' 可以匹配任意字符串。

//输入：s = "cb", p = "?a"
//输出：false
//解释：'?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。

//提示：
//	0 <= s.length, p.length <= 2000
//	s 仅由小写英文字母组成
//	p 仅由小写英文字母、'?' 或 '*' 组成

func isMatch(s string, p string) bool {
	m, n, dp := len(s), len(p), make([][]bool, len(s)+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		if p[i] == '*' {
			dp[0][i+1] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else {
				dp[i][j] = (p[j-1] == '?' || p[j-1] == s[i-1]) && dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func backtrace(s string, p string) bool {
	cnt, idxMap := 0, make(map[int]int)
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] != '*' {
			cnt++
		}
		idxMap[i] = cnt
	}
	return match(s, p, 0, 0, idxMap)
}

func match(s string, p string, si, pi int, idxMap map[int]int) bool {
	if si == len(s) && pi == len(p) {
		return true
	}
	if si != len(s) && pi == len(p) {
		return false
	}
	if p[pi] == '*' {
		for pi = pi + 1; pi < len(p) && p[pi] == '*'; {
			pi++
		}
		if pi == len(p) {
			return true
		}
		for i := len(s) - idxMap[pi]; i >= si; i-- {
			if (p[pi] == '?' || p[pi] == s[i]) && match(s, p, i+1, pi+1, idxMap) {
				return true
			}
		}
		return false
	} else if si < len(s) && (p[pi] == '?' || p[pi] == s[si]) {
		return match(s, p, si+1, pi+1, idxMap)
	}
	return false
}
