package _49

//https://leetcode.cn/problems/group-anagrams/
//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

//示例 1:
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

//输入: strs = [""]
//输出: [[""]]

//输入: strs = ["a"]
//输出: [["a"]]

//提示：
//	1 <= strs.length <= 10^4
//	0 <= strs[i].length <= 100
//	strs[i] 仅包含小写字母

func groupAnagrams(strs []string) [][]string {
	ans, cntMap := make([][]string, 0), make(map[[26]byte][]string)
	for _, str := range strs {
		cnt := [26]byte{}
		for i := 0; i < len(str); i++ {
			cnt[str[i]-'a']++
		}
		cntMap[cnt] = append(cntMap[cnt], str)
	}
	for _, val := range cntMap {
		ans = append(ans, val)
	}
	return ans
}
