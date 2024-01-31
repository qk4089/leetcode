package _140

import "strings"

//https://leetcode.cn/problems/word-break-ii/
//给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。
//注意：词典中的同一个单词可能在分段中被重复使用多次。

//示例 1：

//输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
//输出:["cats and dog","cat sand dog"]

//输入:s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
//输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
//解释: 注意你可以重复使用字典中的单词。

//输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
//输出:[]

//提示：
//	1 <= s.length <= 20
//	1 <= wordDict.length <= 1000
//	1 <= wordDict[i].length <= 10
//	s 和 wordDict[i] 仅有小写英文字母组成
//	wordDict 中所有字符串都 不同

func wordBreak(s string, wordDict []string) []string {
	ans, wdMap := make([]string, 0), make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wdMap[wordDict[i]] = true
	}
	var backtrace func(s string, start int, segment []string, wordDict map[string]bool)
	backtrace = func(s string, start int, segment []string, wordDict map[string]bool) {
		if start == len(s) {
			if len(segment) > 0 {
				ans = append(ans, strings.Join(segment, " "))
			}
			return
		}
		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if wdMap[str] {
				segment = append(segment, str)
				backtrace(s, i+1, segment, wdMap)
				segment = segment[:len(segment)-1]
			}
		}
	}
	backtrace(s, 0, []string{}, wdMap)
	return ans
}
