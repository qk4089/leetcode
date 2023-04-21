package _567

//给你两个字符串s1和s2，写一个函数来判断s2是否包含s1的排列。如果是，返回true；否则，返回false。
//换句话说，s1的排列之一是s2的子串。

//提示
//	1<=s1.length,s2.length<=10^4
//	s1和s2仅包含小写字母

//示例1：
//输入：s1="ab" s2="eidbaooo"
//输出：true
//解释：s2包含s1的排列之一("ba").

//输入：s1="ab" s2="eidboaoo"
//输出：false

func checkInclusion(s1 string, s2 string) bool {
	tMap, wMap := make(map[byte]int), make(map[byte]int)
	for _, c := range []byte(s1) {
		increase(tMap, c)
	}
	length, valid, left, right, bs := len(s1), 0, 0, 0, []byte(s2)
	for right < len(bs) {
		if _, ok := tMap[bs[right]]; ok {
			increase(wMap, bs[right])
			if wMap[bs[right]] == tMap[bs[right]] {
				valid++
			}
		}
		right++
		for right-left >= length {
			if valid == len(tMap) {
				return true
			}
			if _, ok := tMap[bs[left]]; ok {
				if wMap[bs[left]] == tMap[bs[left]] {
					valid--
				}
				wMap[bs[left]] = wMap[bs[left]] - 1
			}
			left++
		}
	}
	return false
}

func increase(m map[byte]int, key byte) {
	if v, ok := m[key]; ok {
		m[key] = v + 1
	} else {
		m[key] = 1
	}
}
