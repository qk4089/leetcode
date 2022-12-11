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
	max, left, right, wMap := 0, 0, 0, make(map[byte]int)
	for ; right < len(s); right++ {
		if v, exist := wMap[s[right]]; exist {
			if max < right-left {
				max = right - left
			}
			if v >= left {
				left = v + 1
			}
		}
		wMap[s[right]] = right
	}
	if max < right-left {
		return right - left
	}
	return max
}
