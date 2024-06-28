package _931

//https://leetcode.cn/problems/minimum-falling-path-sum/
//给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列
//（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是
//(row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。

//示例 1：
//输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
//输出：13
//解释：如图所示，为和最小的两条下降路径

//输入：matrix = [[-19,57],[-40,-5]]
//输出：-59
//解释：如图所示，为和最小的下降路径

//提示：
//	n == matrix.length == matrix[i].length
//	1 <= n <= 100
//	-100 <= matrix[i][j] <= 100

func minFallingPathSum(matrix [][]int) int {
	m, n, result, dp := len(matrix), len(matrix[0]), 10000, make([][]int, len(matrix))
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i], result = matrix[0][i], min(result, matrix[0][i])
	}
	for i := 1; i < m; i++ {
		result = 10000
		for j := 0; j < n; j++ {
			lu, up, ru := 10000, dp[i-1][j], 10000
			if j > 0 {
				lu = min(lu, dp[i-1][j-1])
			}
			if j < n-1 {
				ru = min(ru, dp[i-1][j+1])
			}
			dp[i][j] = min(up+matrix[i][j], min(lu+matrix[i][j], ru+matrix[i][j]))
			result = min(result, dp[i][j])
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
