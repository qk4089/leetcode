package _58

//https://leetcode.cn/problems/length-of-last-word/
//给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。
//单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。

// 示例 1：
//输入：s = "Hello World"
//输出：5
//解释：最后一个单词是“World”，长度为5。

//输入：s = "   fly me   to   the moon  "
//输出：4
//解释：最后一个单词是“moon”，长度为4。

//输入：s = "luffy is still joyboy"
//输出：6
//解释：最后一个单词是长度为6的“joyboy”。

//提示：
//	1 <= s.length <= 10⁴
//	s 仅有英文字母和空格 ' ' 组成
//	s 中至少存在一个单词

func lengthOfLastWord(s string) int {
	idx, ans := len(s)-1, 0
	for ; idx >= 0; idx-- {
		if s[idx] != ' ' {
			break
		}
	}
	for ; idx >= 0; idx-- {
		if s[idx] == ' ' {
			break
		}
		ans++
	}
	return ans
}
