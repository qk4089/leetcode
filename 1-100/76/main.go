package _76

import "math"

//给你一个字符串s、一个字符串t。返回s中涵盖t所有字符的最小子串。如果s中不存在涵盖t所有字符的子串，则返回空字符串""。
//注意
//	对于t中重复字符，我们寻找的子字符串中该字符数量必须不少于t中该字符数量。
//	如果s中存在这样的子串，我们保证它是唯一的答案。

//提示：
//	1<=s.length,t.length<=10^5
//	s和t由英文字母组成

//示例
//输入：s="ADOBECODEBANC",t="ABC"
//输出："BANC"

//输入：s="a",t="a"
//输出："a"

//输入:s="a",t="aa"
//输出:""
//解释:t中两个字符'a'均应包含在s的子串中， 因此没有符合条件的子字符串，返回空字符串。

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	minLeft, minRight, minLength, valid, left, right, sMap := 0, 0, math.MaxInt, 0, 0, 0, make(map[byte]int)
	for right < len(s) {
		if _, ok := tMap[s[right]]; ok {
			sMap[s[right]]++
			if sMap[s[right]] == tMap[s[right]] {
				valid++
			}
		}
		right++
		for valid >= len(tMap) {
			if right-left < minLength {
				minLeft, minRight, minLength = left, right, right-left
			}
			if _, ok := tMap[s[left]]; ok {
				if sMap[s[left]] == tMap[s[left]] {
					valid--
				}
				sMap[s[left]]--
			}
			left++
		}
	}
	if minLength != math.MaxInt {
		return s[minLeft:minRight]
	} else {
		return ""
	}
}
