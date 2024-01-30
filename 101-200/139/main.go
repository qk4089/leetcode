package _139

//https://leetcode.cn/problems/word-break/
//给你一个字符串s和一个字符串列表wordDict作为字典。请你判断是否可以利用字典中出现的单词拼接出s。
//注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

//示例
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//注意，你可以重复使用字典中的单词。

// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出: false

//提示：
//	1 <= s.length <= 300
//	1 <= wordDict.length <= 1000
//	1 <= wordDict[i].length <= 20
//	s 和 wordDict[i] 仅有小写英文字母组成
//	wordDict 中的所有字符串 互不相同

func wordBreak(s string, wordDict []string) bool {
	wdMap := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wdMap[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wdMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// timeout
func backtrace(s string, wordDict map[string]int) bool {
	if len(s) == 0 {
		return true
	}
	for key, val := range wordDict {
		if len(s) >= val && s[:val] == key && backtrace(s[val:], wordDict) {
			return true
		}
	}
	return false
}
