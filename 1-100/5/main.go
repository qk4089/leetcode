package _5

//给你一个字符串s，找到s中最长的回文子串。

//提示：
//	1<=s.length<=1000
//	s仅由数字和英文字母组成

//示例
//输入：s="babad"
//输出："bab"
//解释："aba"同样是符合题意的答案。
//
//输入：s="cbbd"
//输出："bb"

func longestPalindrome(s string) string {
	result := string(s[0])
	for i := 0; i < len(s)-1; i++ {
		left := getPalindrome(s, i, i)
		right := getPalindrome(s, i, i+1)
		result = max(result, max(left, right))
	}
	return result
}

func getPalindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func max(x, y string) string {
	if len(x) > len(y) {
		return x
	}
	return y
}
