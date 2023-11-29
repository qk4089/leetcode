package _14

//https://leetcode.cn/problems/longest-common-prefix/
//编写一个函数来查找字符串数组中的最长公共前缀。

//如果不存在公共前缀，返回空字符串 ""。

//示例 1：
//输入：strs = ["flower","flow","flight"]
//输出："fl"

//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。

//提示：
//	1 <= strs.length <= 200
//	0 <= strs[i].length <= 200
//	strs[i] 仅由小写英文字母组成

func longestCommonPrefix(strs []string) string {
	var tmp byte
	ans := make([]byte, 0)
	for idx := 0; idx == 0 || tmp > 0; idx++ {
		if idx >= len(strs[0]) {
			break
		}
		tmp = strs[0][idx]
		for i := 1; i < len(strs); i++ {
			if idx >= len(strs[i]) || tmp != strs[i][idx] {
				tmp = 0
				break
			}
		}
		if tmp > 0 {
			ans = append(ans, tmp)
		}
	}
	return string(ans)
}
