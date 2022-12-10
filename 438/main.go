package _438

//给定两个字符串s和p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//异位词指由相同字母重排列形成的字符串（包括相同的字符串）。

//示例
//输入:s="cbaebabacd",p="abc"
//输出:[0,6]
//解释:
//起始索引等于0的子串是"cba",它是"abc"的异位词。
//起始索引等于6的子串是"bac",它是"abc"的异位词。

//输入:s="abab",p="ab"
//输出:[0,1,2]
//解释:
//起始索引等于0的子串是"ab",它是"ab"的异位词。
//起始索引等于1的子串是"ba",它是"ab"的异位词。
//起始索引等于2的子串是"ab",它是"ab"的异位词。

//提示:
//	1<=s.length,p.length<=3*10^4
//	s和p仅包含小写字母

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	if len(p) > len(s) {
		return result
	}
	valid, left, right, wMap, nMap := 0, 0, 0, make(map[byte]int), make(map[byte]int)
	for _, v := range []byte(p) {
		nMap[v]++
	}
	for right < len(s) {
		if v, ok := nMap[s[right]]; ok {
			wMap[s[right]] = wMap[s[right]] + 1
			if v == wMap[s[right]] {
				valid++
			}
		}
		right++
		for right-left >= len(p) {
			if valid == len(nMap) {
				result = append(result, left)
			}
			if v, ok := nMap[s[left]]; ok {
				if v == wMap[s[left]] {
					valid--
				}
				wMap[s[left]]--
			}
			left++
		}
	}
	return result
}
