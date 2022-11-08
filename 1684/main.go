package _1684

//给你一个由不同字符组成的字符串allowed和一个字符串数组words。如果一个字符串的每一个字符都在allowed中，就称这个字符串是一致字符串。
//请你返回words数组中一致字符串的数目。
//提示：
//	1 <= words.length <= 104
//	1 <= allowed.length <= 26
//	1 <= words[i].length <= 10
//	allowed中的字符互不相同
//	words[i]和allowed只包含小写英文字母

//示例
//输入：allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
//输出：2
//解释：字符串 "aaab" 和 "baa" 都是一致字符串，因为它们只包含字符 'a' 和 'b'

//输入：allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
//输出：7

//输入：allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
//输出：4

// hashtable
func main(allowed string, words []string) int {
	var aMap = make(map[byte]byte, len(allowed))
	for _, val := range []byte(allowed) {
		aMap[val] = 'y'
	}
	count := len(words)
	for _, word := range words {
		for _, w := range []byte(word) {
			if _, ok := aMap[w]; !ok {
				count--
				break
			}
		}
	}
	return count
}

// 使用数组模拟 hashtable

func main_2(allowed string, words []string) int {
	var arr = make([]byte, 26)
	for _, val := range []byte(allowed) {
		arr[val-'a'] = '1'
	}
	count := len(words)
	for _, word := range words {
		for _, w := range []byte(word) {
			if arr[w-'a'] != '1' {
				count--
				break
			}
		}
	}
	return count
}
