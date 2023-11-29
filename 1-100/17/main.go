package _17

//https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

//示例 1：
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

//输入：digits = ""
//输出：[]

//输入：digits = "2"
//输出：["a","b","c"]

//提示：
//	0 <= digits.length <= 4
//	digits[i] 是范围 ['2', '9'] 的一个数字。

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	bytes, byteMap := make([][]byte, 0), map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'}}
	for i := 0; i < len(digits); i++ {
		bytes = append(bytes, byteMap[digits[i]])
	}
	return convert(link(bytes))
}

func link(bytes [][]byte) [][]byte {
	if len(bytes) == 1 {
		tmp := make([][]byte, 0)
		for _, b := range bytes[0] {
			tmp = append(tmp, []byte{b})
		}
		return tmp
	}
	result, linkedStr := make([][]byte, 0), link(bytes[1:])
	for _, b := range bytes[0] {
		for _, str := range linkedStr {
			result = append(result, append([]byte{b}, str...))
		}
	}
	return result
}

func convert(bytes [][]byte) []string {
	result := make([]string, 0)
	for _, b := range bytes {
		result = append(result, string(b))
	}
	return result
}
