package _48

//给定一个n×n的二维矩阵matrix表示一个图像。请你将图像顺时针旋转90度。
//你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

//提示：
//	n==matrix.length==matrix[i].length
//	1<=n<=20
//	-1000<=matrix[i][j]<=1000

//示例
//输入：matrix=[[1,2,3],[4,5,6],[7,8,9]]
//输出：[[7,4,1],[8,5,2],[9,6,3]]

//输入：matrix=[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j, k := 0, len(matrix)-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
