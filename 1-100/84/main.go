package _84

//https://leetcode.cn/problems/largest-rectangle-in-histogram/
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//求在该柱状图中，能够勾勒出来的矩形的最大面积。

//示例 1:

//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10

//输入： heights = [2,4]
//输出： 4

//提示：
//	1 <= heights.length <=10^5
//	0 <= heights[i] <= 10^4

func largestRectangleArea(heights []int) int {
	ans, queue := 0, make([]int, 0)
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		if len(queue) == 0 || heights[queue[len(queue)-1]] <= heights[i] {
			queue = append(queue, i)
			continue
		}
		for idx := len(queue) - 1; idx >= 0 && heights[queue[idx]] > heights[i]; idx-- {
			if idx > 0 && heights[queue[idx]] == heights[queue[idx-1]] {
				continue
			}
			if idx == 0 {
				ans = max(ans, heights[queue[idx]]*i)
			} else {
				ans = max(ans, heights[queue[idx]]*(i-queue[idx-1]-1))
			}
			queue = queue[:idx]
		}
		queue = append(queue, i)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
