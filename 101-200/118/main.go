package _118

//https://leetcode.cn/problems/pascals-triangle/
//给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//在「杨辉三角」中，每个数是它左上方和右上方的数的和。

//示例 1:
//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

//输入: numRows = 1
//输出: [[1]]

//提示:
//	1 <= numRows <= 30

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	if numRows >= 1 {
		ans[0] = []int{1}
	}
	for i := 1; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans[i][i] = 1
	}
	return ans
}
