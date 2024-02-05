package _74

//https://leetcode.cn/problems/search-a-2d-matrix/
//你一个满足下述两条属性的 m x n 整数矩阵：
//	每行中的整数从左到右按非严格递增顺序排列。
//	每行的第一个整数大于前一行的最后一个整数。
//给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

//示例 1：
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true

//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false

//提示：
//	m == matrix.length
//	n == matrix[i].length
//	1 <= m, n <= 100
//	-10^4 <= matrix[i][j], target <= 10^4

func searchMatrix(matrix [][]int, target int) bool {
	if exist, row := bSearchCol(matrix, target); !exist {
		return row != -1 && bSearchRow(matrix[row], target)
	}
	return true
}

func bSearchCol(matrix [][]int, target int) (bool, int) {
	start, end := 0, len(matrix)-1
	if matrix[start][0] == target || matrix[end][0] == target {
		return true, 0
	} else if matrix[start][0] > target {
		return false, -1
	} else if matrix[end][0] < target {
		return false, end
	}
	for start < end {
		mid := start + (end-start)>>1
		if matrix[mid][0] == target {
			return true, 0
		} else if matrix[mid][0] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return false, start - 1
}

func bSearchRow(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	if nums[end] == target {
		return true
	} else if nums[end] < target {
		return false
	}
	for start <= end {
		mid := start + (end-start)>>1
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
