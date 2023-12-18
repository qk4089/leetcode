package _93

import (
	"strconv"
	"strings"
)

//https://leetcode.cn/problems/restore-ip-addresses/
//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
//给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。
//你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

//示例 1：
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]

//输入：s = "0000"
//输出：["0.0.0.0"]

//输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

//提示：
//	1 <= s.length <= 20
//	s 仅由数字组成

func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	backtrace(0, 1, s, []string{}, &ans)
	return ans
}

func backtrace(segStart, segIdx int, s string, trace []string, ans *[]string) {
	if len(trace) == 4 {
		if segStart == len(s) {
			*ans = append(*ans, strings.Join(trace, "."))
		}
		return
	}
	for i := 1; segStart+i <= len(s) && i <= 3; i++ {
		if isValidSegment(segStart, segStart+i, segIdx, s) {
			trace = append(trace, s[segStart:segStart+i])
			backtrace(segStart+i, segIdx+1, s, trace, ans)
			trace = trace[:len(trace)-1]
		}
	}
}

func isValidSegment(segStart, segEnd, segIdx int, s string) bool {
	if segIdx < 4 && len(s)-segEnd > 3*(4-segIdx) {
		return false
	}
	tmpStr := s[segStart:segEnd]
	if len(tmpStr) > 1 && tmpStr[0] == '0' {
		return false
	}
	tmp, _ := strconv.Atoi(tmpStr)
	return tmp <= 255
}
