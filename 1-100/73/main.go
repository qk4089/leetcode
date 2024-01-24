package _73

//https://leetcode.cn/problems/set-matrix-zeroes/
//给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

//示例 1：
//输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
//输出：[[1,0,1],[0,0,0],[1,0,1]]

//输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
//输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

//提示：
//	m == matrix.length
//	n == matrix[0].length
//	1 <= m, n <= 200
//-	2^31 <= matrix[i][j] <= 2^31 - 1

//进阶：
//	一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
//	一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
//	你能想出一个仅使用常量空间的解决方案吗？

func setZeroes(matrix [][]int) {
	m, n, zero := len(matrix), len(matrix[0]), make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zero = append(zero, []int{i, j})
			}
		}
	}
	for i := 0; i < len(zero); i++ {
		x, y := zero[i][0], zero[i][1]
		matrix[x][y] = 0
		for lx := x - 1; lx >= 0; lx-- {
			matrix[lx][y] = 0
		}
		for rx := x + 1; rx < m; rx++ {
			matrix[rx][y] = 0
		}
		for uy := y - 1; uy >= 0; uy-- {
			matrix[x][uy] = 0
		}
		for dy := y + 1; dy < n; dy++ {
			matrix[x][dy] = 0
		}
	}
}
