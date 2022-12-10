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
	if len(s) == 1 {
		return s
	}
	start, end := -1, -1
	for point := 0; point < len(s)-1 && len(s)-point > (end-start)/2; point++ {
		if _start, _end, ok := getMaxPalindrome(s, point, point); ok {
			start, end = max(start, end, _start, _end)
		}
		if _start, _end, ok := getMaxPalindrome(s, point, point+1); ok {
			start, end = max(start, end, _start, _end)
		}
	}
	return string(s[start : end+1])
}

func getMaxPalindrome(s string, p0 int, p1 int) (int, int, bool) {
	flag := false
	for p0 >= 0 && p1 < len(s) && s[p0] == s[p1] {
		flag = true
		p0, p1 = p0-1, p1+1
	}
	return p0 + 1, p1 - 1, flag
}

func max(s1 int, e1 int, s2 int, e2 int) (int, int) {
	if e1-s1 > e2-s2 {
		return s1, e1
	}
	return s2, e2
}
