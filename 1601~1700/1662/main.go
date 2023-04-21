package _1662

//给你两个字符串数组word1和word2。如果两个数组表示的字符串相同，返回true；否则，返回false。
//
//数组表示的字符串是由数组中的所有元素按顺序连接形成的字符串。
//提示：
//	1 <= word1.length, word2.length <= 103
//	1 <= word1[i].length, word2[i].length <= 103
//	1 <= sum(word1[i].length), sum(word2[i].length) <= 103
//	word1[i] 和 word2[i] 由小写字母组成

//示例
//输入：word1 = ["ab", "c"], word2 = ["a", "bc"]
//输出：true
//解释：
//word1 表示的字符串为 "ab" + "c" -> "abc"
//word2 表示的字符串为 "a" + "bc" -> "abc"
//两个字符串相同，返回 true

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	p1, p2, sum1, sum2 := 0, 0, 0, 0
	l1, l2 := len(word1), len(word2)
	for p1 < l1 {
		sum1, p1 = sum1+len(word1[p1]), p1+1
	}
	for p2 < l2 {
		sum2, p2 = sum2+len(word2[p2]), p2+1
	}
	if sum1 != sum2 {
		return false
	}
	index := 0
	p1, p2 = 0, 0
	s1, s2 := []byte(word1[p1]), []byte(word2[p2])
LABEL:
	for index < sum1 {
		if index >= len(s1) {
			p1++
			s1 = append(s1, []byte(word1[p1])...)
			goto LABEL
		}
		if index >= len(s2) {
			p2++
			s2 = append(s2, []byte(word2[p2])...)
			goto LABEL
		}
		if s1[index] != s2[index] {
			return false
		} else {
			index++
		}
	}
	return true
}
