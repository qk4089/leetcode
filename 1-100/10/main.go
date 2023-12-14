package _10

//https://leetcode.cn/problems/regular-expression-matching/
//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//	'.' 匹配任意单个字符
//	'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖 整个字符串s的，而不是部分字符串。

//示例 1：
//输入：s = "aa", p = "a"
//输出：false
//解释："a" 无法匹配 "aa" 整个字符串。

//输入：s = "aa", p = "a*"
//输出：true
//解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

//输入：s = "ab", p = ".*"
//输出：true
//解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

//提示：
//	1 <= s.length <= 20
//	1 <= p.length <= 20
//	s 只包含从 a-z 的小写字母。
//	p 只包含从 a-z 的小写字母，以及字符 . 和 *。
//	保证每次出现字符 * 时，前面都匹配到有效的字符

func isMatch(s string, p string) bool {
	if len(s)+len(p) == 0 {
		return true
	}
	if len(s) != 0 && len(p) == 0 {
		return false
	}
	if len(p) > 1 && p[1] == '*' {
		if isMatch(s, p[2:]) {
			return true
		}
		if len(s) > 0 && (p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p) {
			return true
		}
	} else if len(s) > 0 && (p[0] == '.' || p[0] == s[0]) {
		return isMatch(s[1:], p[1:])
	}
	return false
}
