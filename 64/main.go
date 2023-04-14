package _64

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//说明：每次只能向下或者向右移动一步。

//示例
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//  1   3   1
//  1   5   1
//  4   2   1

//输入：grid = [[1,2,3],[4,5,6]]
//输出：12

//提示：
//	m == grid.length
//	n == grid[i].length
//	1 <= m, n <= 200
//	0 <= grid[i][j] <= 100

func minPathSum(grid [][]int) int {
	m, n, result := len(grid), len(grid[0]), make([][]int, len(grid))
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	result[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		result[0][i] = result[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		result[i][0] = result[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			result[i][j] = grid[i][j] + min(result[i-1][j], result[i][j-1])
		}
	}
	return result[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
