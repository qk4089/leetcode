package _90

import "sort"

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

//示例 1：
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]

//输入：nums = [0]
//输出：[[],[0]]

//提示：
//	1 <= nums.length <= 10
//	-10 <= nums[i] <= 10

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{{}, nums}
	for i := 1; i < len(nums); i++ {
		backtrace(0, i, []int{}, nums, &ans)
	}
	return ans
}

func backtrace(start, cnt int, trace, nums []int, ans *[][]int) {
	if len(trace) == cnt {
		*ans = append(*ans, append([]int{}, trace...))
		return
	}
	for i := start; i < len(nums); i++ {
		if start != i && nums[i] == nums[i-1] {
			continue
		}
		trace = append(trace, nums[i])
		backtrace(i+1, cnt, trace, nums, ans)
		trace = trace[:len(trace)-1]
	}
}
