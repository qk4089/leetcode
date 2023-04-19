package _62

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//问总共有多少条不同的路径？

//示例
//输入：m = 3, n = 7
//输出：28

//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下

//输入：m = 7, n = 3
//输出：28

//输入：m = 3, n = 3
//输出：6

//提示：
//	1 <= m, n <= 100
//	题目数据保证答案小于等于 2 * 10^9

func uniquePaths(m int, n int) int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		result[i][0] = 1
	}
	for i := 0; i < n; i++ {
		result[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}
	return result[m-1][n-1]
}
