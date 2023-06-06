package _304

//给定一个二维矩阵matrix，以下类型的多个请求：
//计算其子矩形范围内元素的总和，该子矩阵的左上角为(row1,col1)，右下角为(row2,col2)。
//实现NumMatrix类：
//	NumMatrix(int[][]matrix)给定整数矩阵matrix进行初始化
//	int sumRegion(int row1,int col1,int row2,int col2)返回左上角(row1,col1)、右下角(row2,col2)所描述的子矩阵的元素总和。

//提示
//	m==matrix.length
//	n==matrix[i].length
//	1<=m,n<=200
//	-105<=matrix[i][j]<=105
//	0<=row1<=row2<m
//	0<=col1<=col2<n
//	最多调用104次sumRegion方法

//示例1
//输入:
//	["NumMatrix","sumRegion","sumRegion","sumRegion"]
//	[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
//输出:
//	[null,8,11,12]
//解释:
//	NumMatrix numMatrix=newNumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
//	numMatrix.sumRegion(2,1,4,3);//return 8(红色矩形框的元素总和)
//	numMatrix.sumRegion(1,1,2,2);//return 11(绿色矩形框的元素总和)
//	numMatrix.sumRegion(1,2,2,4);//return 12(蓝色矩形框的元素总和)

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rL, cL := len(matrix)+1, len(matrix[0])+1
	result := make([][]int, rL)
	for i := range result {
		result[i] = make([]int, cL)
	}
	for i := 1; i <= rL; i++ {
		for j := 1; j <= cL; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1] + matrix[i-1][j-1] - result[i-1][j-1]
		}
	}
	return NumMatrix{sum: result}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row2+1][col1] - this.sum[row1][col2+1] + this.sum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
