package _3

//给定一个字符串s，请你找出其中不含有重复字符的最长子串的长度。

//提示：
//	0<=s.length<=5*10^4
//	s由英文字母、数字、符号和空格组成

//示例
//输入:s="abcabcbb"
//输出:3
//解释:因为无重复字符的最长子串是"abc"，所以其长度为3。

//输入:s="bbbbb"
//输出:1
//解释:因为无重复字符的最长子串是"b"，所以其长度为1。

//输入:s="pwwkew"
//输出:3
//解释:因为无重复字符的最长子串是"wke"，所以其长度为3。
//请注意，你的答案必须是子串的长度，"pwke"是一个子序列，不是子串。

func lengthOfLongestSubstring(s string) int {
	maxV, left, right, window := 0, 0, 0, make(map[byte]int)
	for ; right < len(s); right++ {
		if idx, exist := window[s[right]]; exist {
			left, maxV = max(left, idx+1), max(maxV, right-left)
		}
		window[s[right]] = right
	}
	return max(maxV, right-left)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
