package _57

//https://leetcode.cn/problems/insert-interval/
//给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
//在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

//示例 1：
//
//输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
//输出：[[1,5],[6,9]]

//输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出：[[1,2],[3,10],[12,16]]
//解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

//输入：intervals = [], newInterval = [5,7]
//输出：[[5,7]]

//输入：intervals = [[1,5]], newInterval = [2,3]
//输出：[[1,5]]

//输入：intervals = [[1,5]], newInterval = [2,7]
//输出：[[1,7]]

//提示：
//	0 <= intervals.length <= 10^4
//	intervals[i].length == 2
//	0 <= intervals[i][0] <= intervals[i][1] <= 10^5
//	intervals 根据 intervals[i][0] 按 升序 排列
//	newInterval.length == 2
//	0 <= newInterval[0] <= newInterval[1] <= 10^5

func insert(intervals [][]int, newInterval []int) [][]int {
	idx, ans := 0, make([][]int, 0)
	for ; idx < len(intervals); idx++ {
		if newInterval[0] <= intervals[idx][1] {
			break
		}
		ans = append(ans, intervals[idx])
	}
	if idx == len(intervals) {
		return append(ans, newInterval)
	}
	if newInterval[1] >= intervals[idx][0] {
		start, end := min(intervals[idx][0], newInterval[0]), max(intervals[idx][1], newInterval[1])
		for ; idx < len(intervals) && newInterval[1] >= intervals[idx][0]; idx++ {
			end = max(end, intervals[idx][1])
		}
		ans = append(ans, []int{start, end})
	} else {
		ans = append(ans, newInterval)
	}
	return append(ans, intervals[idx:]...)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
