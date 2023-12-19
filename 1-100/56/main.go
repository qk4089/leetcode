package _56

//https://leetcode.cn/problems/merge-intervals/
//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
//请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

//示例 1：
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

//提示：
//	1 <= intervals.length <= 10^4
//	intervals[i].length == 2
//	0 <= starti <= endi <= 10^4

func merge(intervals [][]int) [][]int {
	maxLen := 0
	for _, interval := range intervals {
		if interval[1] > maxLen {
			maxLen = interval[1]
		}
	}
	ans, diff := make([][]int, 0), make([]int, maxLen*2+2)
	for _, interval := range intervals {
		diff[interval[0]*2]++
		diff[interval[1]*2+1]--
	}
	start := -1
	if diff[0] > 0 {
		start = 0
	}
	for i := 1; i < len(diff); i++ {
		diff[i] += diff[i-1]
		if diff[i] == 0 && start >= 0 {
			start, ans = -1, append(ans, []int{start / 2, (i - 1) / 2})
		} else if diff[i] > 0 && start < 0 {
			start = i
		}
	}
	return ans
}
