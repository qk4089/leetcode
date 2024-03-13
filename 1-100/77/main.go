package _77

//https://leetcode.cn/problems/combinations/
//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。你可以按 任何顺序 返回答案。

//示例 1：
//输入：n = 4, k = 2
//输出：
//[
//	[2,4],
//	[3,4],
//	[2,3],
//	[1,2],
//	[1,3],
//	[1,4],
//]

//输入：n = 1, k = 1
//输出：[[1]]

//提示：
//	1 <= n <= 20
//	1 <= k <= n

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	var backtrace func(int, []int)
	backtrace = func(start int, trace []int) {
		if len(trace) == k {
			ans = append(ans, append([]int{}, trace...))
			return
		}
		for i := start; i <= n; i++ {
			trace = append(trace, i)
			backtrace(i+1, trace)
			trace = trace[:len(trace)-1]
		}
	}
	backtrace(1, make([]int, 0))
	return ans
}
