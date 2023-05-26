package _151

//给你一个字符串s，请你反转字符串中单词的顺序。
//单词是由非空格字符组成的字符串。s中使用至少一个空格将字符串中的单词分隔开。
//返回单词顺序颠倒且单词之间用单个空格连接的结果字符串。
//注意：输入字符串s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

//提示：
//	1<=s.length<=104
//	s包含英文大小写字母、数字和空格''
//	s中至少存在一个单词

//示例
//输入：s="the sky is blue"
//输出："blue is sky the"

//输入：s="  hello world  "
//输出："world hello"
//解释：反转后的字符串中不能存在前导空格和尾随空格。

//输入：s="a good   example"
//输出："example good a"
//解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。

//进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用O(1)额外空间复杂度的原地解法。

func reverseWords(s string) string {
	start, flag, sc, b := 0, false, 0, []byte(s)
	reverse(b, 0, len(b)-1)
	i := 0
	for ; i < len(b); i++ {
		if b[i] != ' ' {
			flag = true
		} else if flag {
			reverse(b, start, i-1)
			start, flag = i-sc+1, false
			sc = i - start + 1
		} else {
			sc++
		}
	}
	if flag {
		reverse(b, start, i-1)
		start = i - sc + 1
	}
	return string(b[:start-1])
}

func reverse(s []byte, start, end int) {
	for start < end {
		s[start], s[end], start, end = s[end], s[start], start+1, end-1
	}
}
