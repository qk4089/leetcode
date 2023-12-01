package _51

//https://leetcode.cn/problems/n-queens/
//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。n皇后问题研究的是如何将n个皇后放置在n×n的棋盘上，
//并且使皇后彼此之间不能相互攻击。

//给你一个整数n，返回所有不同的n皇后问题的解决方案。每一种解法包含一个不同的n皇后问题的棋子放置方案，该方案中'Q'和'.'分别代表了皇后和空位。

//示例 1：

//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。

//输入：n = 1
//输出：[["Q"]]

//提示：
//	1 <= n <= 9

func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	backtrace(0, make([]string, n), &ans)
	return ans
}

func backtrace(row int, position []string, result *[][]string) {
	if row >= len(position) {
		*result = append(*result, append([]string{}, position...))
		return
	}
	for i := 0; i < len(position); i++ {
		if isMatch(row, i, position) {
			position[row] = generate(i, len(position))
			backtrace(row+1, position, result)
		}
	}
}

func generate(idx, n int) string {
	str := make([]byte, n)
	for i := 0; i < n; i++ {
		if i == idx {
			str[i] = 'Q'
		} else {
			str[i] = '.'
		}
	}
	return string(str)
}

func isMatch(row, idx int, position []string) bool {
	if row == 0 {
		return true
	}
	for i, cnt := row-1, 1; i >= 0; i, cnt = i-1, cnt+1 {
		str, left, right := position[i], idx-cnt, idx+cnt
		if str[idx] == 'Q' || (left >= 0 && str[left] == 'Q') || (right < len(position) && str[right] == 'Q') {
			return false
		}
	}
	return true
}
