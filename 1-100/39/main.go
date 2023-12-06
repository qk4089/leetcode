package _39

import "sort"

//https://leetcode.cn/problems/combination-sum/
//给你一个无重复元素的整数数组candidates和一个目标整数target，找出candidates中可以使数字和为目标数target的所有不同组合，并以列表形式返回。
//你可以按 任意顺序 返回这些组合。

//candidates中的同一个数字可以无限制重复被选取。如果至少一个数字的被选数量不同，则两种组合是不同的。
//对于给定的输入，保证和为target的不同组合数少于 150 个。

//示例 1：
//输入：candidates = [2,3,6,7], target = 7
//输出：[[2,2,3],[7]]
//解释：
//2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
//7 也是一个候选， 7 = 7 。
//仅有这两种组合。

//输入: candidates = [2,3,5], target = 8
//输出: [[2,2,2,2],[2,3,3],[3,5]]

//输入: candidates = [2], target = 1
//输出: []

//提示：
//	1 <= candidates.length <= 30
//	2 <= candidates[i] <= 40
//	candidates 的所有元素 互不相同
//	1 <= target <= 40

func combinationSum(candidates []int, target int) [][]int {
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
		trace = append(trace, candidates[i])
		if sum+candidates[i] > target {
			break
		}
		backtrace(i, sum+candidates[i], target, trace, candidates, ans)
		trace = trace[:len(trace)-1]
	}
}
