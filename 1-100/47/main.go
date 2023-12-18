package _47

import (
	"sort"
)

//https://leetcode.cn/problems/permutations-ii/
//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

//示例 1：
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
//[1,2,1],
//[2,1,1]]

//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

//提示：
//	1 <= nums.length <= 8
//	-10 <= nums[i] <= 10

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	backtrace(nums, []int{}, make([]bool, len(nums)), &ans)
	return ans
}

func backtrace(nums, trace []int, used []bool, ans *[][]int) {
	if len(trace) == len(nums) {
		*ans = append(*ans, append([]int{}, trace...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] || i != 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		trace, used[i] = append(trace, nums[i]), true
		backtrace(nums, trace, used, ans)
		trace, used[i] = trace[:len(trace)-1], false
	}
}
