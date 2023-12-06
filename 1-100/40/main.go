package _40

import "sort"

//https://leetcode.cn/problems/combination-sum-ii/
//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//candidates 中的每个数字在每个组合中只能使用 一次 。
//注意：解集不能包含重复的组合。

//示例 1:
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出: [[1,1,6], [1,2,5], [1,7], [2,6]]

//输入: candidates = [2,5,2,1,2], target = 5,
//输出: [[1,2,2], [5]]

//提示:
//	1 <= candidates.length <= 100
//	1 <= candidates[i] <= 50
//	1 <= target <= 30

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	backtrace(0, 0, target, make([]int, 0), candidates, &ans)
	return ans
}

func backtrace(start, sum, target int, trace, candidates []int, ans *[][]int) {
	if sum == target {
		*ans = append(*ans, append([]int{}, trace...))
		return
	}
	for i := start; i < len(candidates); i++ {
		if i != start && candidates[i] == candidates[i-1] {
			continue
		}
		trace = append(trace, candidates[i])
		if sum+candidates[i] > target {
			break
		}
		backtrace(i+1, sum+candidates[i], target, trace, candidates, ans)
		trace = trace[:len(trace)-1]
	}
}
