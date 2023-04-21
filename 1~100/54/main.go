package _54

//给你一个m行n列的矩阵matrix，请按照顺时针螺旋顺序，返回矩阵中的所有元素。

//提示：
//	m==matrix.length
//	n==matrix[i].length
//	1<=m,n<=10
//	-100<=matrix[i][j]<=100

//示例
//输入：matrix=[[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]

//输入：matrix=[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]

func spiralOrder(matrix [][]int) []int {
	var result []int
	spiral(&result, matrix, 0, len(matrix[0]), len(matrix))
	return result
}

func spiral(result *[]int, matrix [][]int, startRow, wide, height int) {
	if wide <= 1 || height <= 1 {
		for i := startRow; i < startRow+height; i++ {
			for j := startRow; j < startRow+wide; j++ {
				*result = append(*result, matrix[i][j])
			}
		}
		return
	}
	for up := startRow; up < startRow+wide; up++ {
		*result = append(*result, matrix[startRow][up])
	}
	for right := startRow + 1; right < startRow+height-1; right++ {
		*result = append(*result, matrix[right][startRow+wide-1])
	}
	for down := startRow + wide - 1; down >= startRow; down-- {
		*result = append(*result, matrix[startRow+height-1][down])
	}
	for left := startRow + height - 2; left > startRow; left-- {
		*result = append(*result, matrix[left][startRow])
	}
	spiral(result, matrix, startRow+1, wide-2, height-2)
}
