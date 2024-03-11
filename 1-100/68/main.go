package _68

import "strings"

//https://leetcode.cn/problems/text-justification/
//给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
//你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
//要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

//文本的最后一行应为左对齐，且单词之间不插入额外的空格。

//注意:
//	单词是指由非空格字符组成的字符序列。
//	每个单词的长度大于 0，小于等于 maxWidth。
//	输入单词数组 words 至少包含一个单词。

//示例 1:
//输入: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
//输出:
//[
//	"This    is    an",
//	"example  of text",
//	"justification.  "
//]

//输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
//输出:
//[
//	"What   must   be",
//	"acknowledgment  ",
//	"shall be        "
//]
//解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
//因为最后一行应为左对齐，而不是左右两端对齐。
//第二行同样为左对齐，这是因为这行只包含一个单词。

//输入:words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth = 20
//输出:
//[
//	"Science  is  what we",
//	"understand      well",
//	"enough to explain to",
//	"a  computer.  Art is",
//	"everything  else  we",
//	"do                  "
//]

//提示:
//	1 <= words.length <= 300
//	1 <= words[i].length <= 20
//	words[i] 由小写英文字母和符号组成
//	1 <= maxWidth <= 100
//	words[i].length <= maxWidth

func fullJustify(words []string, maxWidth int) []string {
	tmp, tmpWidth, ans := []string{words[0]}, len(words[0]), make([]string, 0)
	for i := 1; i < len(words); i++ {
		if tmpWidth+len(words[i])+len(tmp)*1 <= maxWidth {
			tmp, tmpWidth = append(tmp, words[i]), tmpWidth+len(words[i])
			continue
		}
		ans, tmp, tmpWidth = append(ans, fill(tmp, maxWidth-tmpWidth)), []string{words[i]}, len(words[i])
	}
	pre := strings.Join(tmp, " ")
	return append(ans, pre+strings.Repeat(" ", maxWidth-len(pre)))
}

func fill(words []string, diffWidth int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", diffWidth)
	}
	if len(words) == 2 {
		return strings.Join(words, strings.Repeat(" ", diffWidth))
	}
	quotient, mod := diffWidth/(len(words)-1), diffWidth%(len(words)-1)
	if mod == 0 {
		return strings.Join(words, strings.Repeat(" ", quotient))
	}
	return strings.Join(words[:mod+1], strings.Repeat(" ", quotient+1)) + strings.Repeat(" ", quotient) + strings.Join(words[mod+1:], strings.Repeat(" ", quotient))
}
