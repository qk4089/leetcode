package _52

//https://leetcode.cn/problems/n-queens-ii/
//n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

//示例 1：
//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。

//输入：n = 1
//输出：1

//提示：
//	1 <= n <= 9

func totalNQueens(n int) int {
	cnt := 0
	backtrace(0, make([]string, n), &cnt)
	return cnt
}

func backtrace(row int, position []string, result *int) {
	if row >= len(position) {
		*result++
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
