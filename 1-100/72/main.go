package _72

//https://leetcode.cn/problems/edit-distance/
//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数。
//你可以对一个单词进行如下三种操作： 插入一个字符、删除一个字符、替换一个字符

//示例 1：
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')

//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')

//提示：
//	0 <= word1.length, word2.length <= 500
//	word1 和 word2 由小写英文字母组成

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := 1; i <= len(word2); i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
