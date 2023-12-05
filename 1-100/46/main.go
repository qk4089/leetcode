package _46

//https://leetcode.cn/problems/permutations/
//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

//示例 1：
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

//输入：nums = [0,1]
//输出：[[0,1],[1,0]]

//输入：nums = [1]
//输出：[[1]]

//提示：
//	1 <= nums.length <= 6
//	-10 <= nums[i] <= 10
//	nums 中的所有整数 互不相同

func permute(nums []int) [][]int {
	ans, track := make([][]int, 0), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		track[i] = 11
	}
	backtrace(0, nums, track, &ans)
	return ans
}

func backtrace(idx int, nums, track []int, ans *[][]int) {
	if idx == len(nums) {
		*ans = append(*ans, append([]int{}, track...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if !contains(track, nums[i]) {
			track[idx] = nums[i]
			backtrace(idx+1, nums, track, ans)
			track[idx] = 11
		}
	}
}

func contains(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		}
	}
	return false
}
