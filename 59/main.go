package _59

//给你一个正整数n，生成一个包含1到n2所有元素，且元素按顺时针顺序螺旋排列的nxn正方形矩阵matrix。

//提示：
//	1<=n<=20

// 示例
// 输入：n=3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
//
// 输入：n=1
// 输出：[[1]]
func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	generate(result, 1, 0, n)
	return result
}

func generate(matrix [][]int, start, index, remain int) {
	if remain <= 0 {
		return
	}
	if remain == 1 {
		matrix[index][index] = start
		return
	}
	for up := index; up < index+remain; up++ {
		matrix[index][up], start = start, start+1
	}
	for right := index + 1; right < index+remain-1; right++ {
		matrix[right][index+remain-1], start = start, start+1
	}
	for down := index + remain - 1; down >= index; down-- {
		matrix[index+remain-1][down], start = start, start+1
	}
	for left := index + remain - 2; left > index; left-- {
		matrix[left][index], start = start, start+1
	}
	generate(matrix, start, index+1, remain-2)
}
