package _240

//https://leetcode.cn/problems/search-a-2d-matrix-ii/
//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//	每行的元素从左到右升序排列。
//	每列的元素从上到下升序排列。

//示例 1：
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
//输出：true

//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
//输出：false

//提示：
//	m == matrix.length
//	n == matrix[i].length
//	1 <= n, m <= 300
//	-10^9 <= matrix[i][j] <= 10^9
//	-10^9 <= target <= 10^9

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if target == matrix[0][0] || target == matrix[0][n-1] || target == matrix[m-1][n-1] {
		return true
	}
	start := 0
	for m >= 1 && start < n {
		if matrix[m-1][start] == target {
			return true
		} else if matrix[m-1][start] < target {
			start++
		} else {
			m--
		}
	}
	return false
}
