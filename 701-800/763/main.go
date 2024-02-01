package _763

//https://leetcode.cn/problems/partition-labels/
//给你一个字符串s。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。
//注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
//返回一个表示每个字符串片段的长度的列表。

//示例 1：
//输入：s = "ababcbacadefegdehijhklij"
//输出：[9,7,8]
//解释：
//划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
//每个字母最多出现在一个片段中。
//像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。

//输入：s = "eccbbbbdec"
//输出：[10]

//提示：
//	1 <= s.length <= 500
//	s 仅由小写英文字母组成

func partitionLabels(s string) []int {
	ans, charMap := make([]int, 0), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}
	point, segementMap := 0, make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if !segementMap[s[i]] {
			segementMap[s[i]] = true
		}
		charMap[s[i]]--
		if charMap[s[i]] == 0 {
			delete(segementMap, s[i])
		}
		if len(segementMap) == 0 {
			point, ans = i+1, append(ans, i-point+1)
		}
	}
	return ans
}
