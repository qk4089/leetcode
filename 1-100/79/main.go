package _79

//https://leetcode.cn/problems/word-search/
//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

//示例 1：
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true

//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
//输出：true

//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
//输出：false

//提示：
//	m == board.length
//	n = board[i].length
//	1 <= m, n <= 6
//	1 <= word.length <= 15
//	board 和 word 仅由大小写英文字母组成

//进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && backtrace(board, word, make(map[[2]int]bool), i, j, 1) {
				return true
			}
		}
	}
	return false
}

func backtrace(board [][]byte, word string, used map[[2]int]bool, i, j, idx int) bool {
	if idx == len(word) {
		return true
	}
	used[[2]int{i, j}] = true
	if i-1 >= 0 && !used[[2]int{i - 1, j}] && board[i-1][j] == word[idx] && backtrace(board, word, used, i-1, j, idx+1) {
		return true
	}
	if i+1 < len(board) && !used[[2]int{i + 1, j}] && board[i+1][j] == word[idx] && backtrace(board, word, used, i+1, j, idx+1) {
		return true
	}
	if j-1 >= 0 && !used[[2]int{i, j - 1}] && board[i][j-1] == word[idx] && backtrace(board, word, used, i, j-1, idx+1) {
		return true
	}
	if j+1 < len(board[0]) && !used[[2]int{i, j + 1}] && board[i][j+1] == word[idx] && backtrace(board, word, used, i, j+1, idx+1) {
		return true
	}
	used[[2]int{i, j}] = false
	return false
}
