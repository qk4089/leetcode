package _76

import (
	"math"
)

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
	wMap, tMap := make(map[byte]int), make(map[byte]int)
	for _, c := range []byte(t) {
		tMap[c]++
	}
	start, length, left, right, valid := 0, math.MaxInt, 0, 0, 0
	for right < len(s) {
		if _, ok := tMap[s[right]]; ok {
			wMap[s[right]]++
			if wMap[s[right]] == tMap[s[right]] {
				valid++
			}
		}
		right++
		for valid == len(tMap) {
			if right-left < length {
				start, length = left, right-left
			}
			if _, ok := tMap[s[left]]; ok {
				if wMap[s[left]] == tMap[s[left]] {
					valid--
				}
				wMap[s[left]] = wMap[s[left]] - 1
			}
			left++
		}
	}
	if length == math.MaxInt {
		return ""
	} else {
		return string(s[start : start+length])
	}
}

func increase(m map[byte]int, key byte) {
	if v, ok := m[key]; ok {
		m[key] = v + 1
	} else {
		m[key] = 1
	}
}
