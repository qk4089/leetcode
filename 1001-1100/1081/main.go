package _1081

//返回 s 字典序最小的子序列，该子序列包含 s 的所有不同字符，且只包含一次。
//注意：该题与 316 https://leetcode.com/problems/remove-duplicate-letters/ 相同

//示例 1：
//输入：s = "bcabc"
//输出："abc"

//输入：s = "cbacdcbc"
//输出："acdb"

//提示：
//	1 <= s.length <= 1000
//	s 由小写英文字母组成

func smallestSubsequence(s string) string {
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
