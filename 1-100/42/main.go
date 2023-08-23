package _42

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

//示例 1：
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

//输入：height = [4,2,0,3,2,5]
//输出：9

//提示：
//	n == height.length
//	1 <= n <= 2 * 10^4
//	0 <= height[i] <= 10^5

func trap(height []int) int {
	idx, tmp, result := 0, 0, 0
	for i := 1; i < len(height); i++ {
		if height[i] >= height[idx] {
			idx, result, tmp = i, result+tmp, 0
		} else {
			tmp += height[idx] - height[i]
		}
	}
	if idx < len(height)-2 {
		tmp = 0
		reverseIdx := len(height) - 1
		for i := len(height) - 2; i >= idx; i-- {
			if height[i] >= height[reverseIdx] {
				reverseIdx, result, tmp = i, result+tmp, 0
			} else {
				tmp += height[reverseIdx] - height[i]
			}
		}
	}
	return result
}
