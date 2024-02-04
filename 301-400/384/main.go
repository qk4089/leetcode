package _384

import "bytes"

//https://leetcode.cn/problems/decode-string
//给定一个经过编码的字符串，返回它解码后的字符串。
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

//示例 1：
//输入：s = "3[a]2[bc]"
//输出："aaabcbc"

//输入：s = "3[a2[c]]"
//输出："accaccacc"

//输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"

//输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"

//提示：
//	1 <= s.length <= 30
//	s 由小写英文字母、数字和方括号 '[]' 组成
//	s 保证是一个 有效 的输入。
//	s 中所有整数的取值范围为 [1, 300]

func decodeString(s string) string {
	var decode func() []byte
	idx, ans := 0, bytes.NewBufferString("")
	decode = func() []byte {
		repeat := int(s[idx] - '0')
		for idx = idx + 1; s[idx] != '['; idx++ {
			repeat = repeat*10 + (int)(s[idx]-'0')
		}
		segment := make([]byte, 0)
		for idx = idx + 1; ; idx++ {
			if s[idx] >= 'a' && s[idx] <= 'z' {
				segment = append(segment, s[idx])
			} else if s[idx] >= '0' && s[idx] <= '9' {
				segment = append(segment, decode()...)
			} else if s[idx] == ']' {
				break
			}
		}
		return bytes.Repeat(segment, repeat)
	}
	for ; idx < len(s); idx++ {
		if s[idx] >= 'a' && s[idx] <= 'z' {
			ans.WriteByte(s[idx])
			continue
		}
		ans.Write(decode())
	}
	return ans.String()
}
