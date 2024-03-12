package _65

import "strings"

//https://leetcode.cn/problems/valid-number/
//有效数字（按顺序）可以分成以下几个部分：
//	一个 小数 或者 整数
// （可选）一个 'e' 或 'E' ，后面跟着一个 整数
//小数（按顺序）可以分成以下几个部分：
// （可选）一个符号字符（'+' 或 '-'）
//	下述格式之一：
//		至少一位数字，后面跟着一个点 '.'
//		至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
//		一个点 '.' ，后面跟着至少一位数字
//整数（按顺序）可以分成以下几个部分：
// （可选）一个符号字符（'+' 或 '-'）
//	至少一位数字

//部分有效数字列举如下：["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]
//部分无效数字列举如下：["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
//给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。

//示例 1：
//输入：s = "0"
//输出：true

//输入：s = "e"
//输出：false

//输入：s = "."
//输出：false

//提示：
//	1 <= s.length <= 20
//	s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，或者点 '.' 。

func isNumber(s string) bool {
	leNum, ueNum := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'Z' && s[i] != 'e' && s[i] != 'E' {
			return false
		}
		if s[i] == 'e' {
			leNum++
		} else if s[i] == 'E' {
			ueNum++
		}
		if leNum+ueNum > 1 {
			return false
		}
	}
	bytes := make([]string, 0)
	if leNum > 0 {
		bytes = strings.Split(s, "e")
	} else {
		bytes = strings.Split(s, "E")
	}
	if len(bytes) == 2 {
		return isValidNumber(bytes[0]) && isValidInteger(bytes[1])
	} else {
		return isValidNumber(bytes[0])
	}
}

func isValidNumber(s string) bool {
	if !strings.Contains(s, ".") {
		return isValidInteger(s)
	}
	s, size := removeSymbol(s)
	if s == "" || size > 1 {
		return false
	}
	nums := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			nums++
		}
	}
	return nums > 0 && nums >= len(s)-1
}

func isValidInteger(s string) bool {
	s, size := removeSymbol(s)
	if s == "" || size > 1 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func removeSymbol(s string) (string, int) {
	idx := 0
	for idx < len(s) && (s[idx] == '-' || s[idx] == '+') {
		idx++
	}
	return s[idx:], idx
}
