package _131

//https://leetcode.cn/problems/palindrome-partitioning/
//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//回文串 是正着读和反着读都一样的字符串。

//示例 1：
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]

//输入：s = "a"
//输出：[["a"]]

//提示：
//	1 <= s.length <= 16
//	s 仅由小写英文字母组成

func partition(s string) [][]string {
	ans, trace, memo, backtrace := make([][]string, 0), make([]string, 0), make(map[string]bool), func(start int) {}
	backtrace = func(start int) {
		if start == len(s) {
			ans = append(ans, append([]string{}, trace...))
			return
		}
		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if _, exist := memo[str]; !exist {
				memo[str] = adjust(str)
			}
			if memo[str] {
				trace = append(trace, str)
				backtrace(i + 1)
				trace = trace[:len(trace)-1]
			}
		}
	}
	backtrace(0)
	return ans
}

func adjust(str string) bool {
	if len(str) == 1 {
		return true
	}
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left, right = left+1, right-1
	}
	return true
}
