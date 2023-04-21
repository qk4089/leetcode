package _1668

//给你一个字符串sequence，如果字符串word连续重复k次形成的字符串是sequence的一个子字符串，那么单词word的重复值为k。
//单词word的最大重复值是单词word在sequence中最大的重复值。如果word不是sequence的子串，那么重复值k为0 。

//给你一个字符串sequence和word，请你返回最大重复值k
//提示：
//	1 <= sequence.length <= 100
//	1 <= word.length <= 100
//	sequence 和 word 都只包含小写英文字母

//示例
//输入：sequence = "ababc", word = "ab"
//输出：2
//解释："abab" 是 "ababc" 的子字符串。

//输入：sequence = "ababc", word = "ba"
//输出：1
//解释："ba" 是 "ababc" 的子字符串，但 "baba" 不是 "ababc" 的子字符串。

func maxRepeating(sequence string, word string) int {
	max, start, index := 0, 0, 0
	sSize, wSize := len(sequence), len(word)
	for index < sSize {
		count := 0
		for {
			for i := 0; i < wSize; i++ {
				if index >= sSize || word[i] != sequence[index] {
					goto LABEL
				}
				index++
			}
			count++
		}
	LABEL:
		if count > max {
			max = count
		}
		if index == sSize {
			return max
		}
		start, index = start+1, start+1
	}
	return max
}
