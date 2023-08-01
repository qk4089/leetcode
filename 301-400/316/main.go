package _316

//给你一个字符串s，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

//示例 1：
//
//输入：s = "bcabc"
//输出："abc"

//输入：s = "cbacdcbc"
//输出："acdb"

//提示：
//	1 <= s.length <= 10^4
//	s由小写英文字母组成

func removeDuplicateLetters(s string) string {
	cMap, iMap := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		cMap[s[i]]++
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if iMap[s[i]] == 0 {
			for len(stack) > 0 && s[i] < stack[len(stack)-1] && cMap[stack[len(stack)-1]] > 0 {
				iMap[stack[len(stack)-1]]--
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, s[i])
			iMap[s[i]]++
		}
		cMap[s[i]]--
	}
	return string(stack[:])
}
