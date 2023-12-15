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
	ans := make([][]int, 0)
	backtrace(nums, []int{}, make(map[int]bool), &ans)
	return ans
}

func backtrace(nums, track []int, used map[int]bool, ans *[][]int) {
	if len(track) == len(nums) {
		*ans = append(*ans, append([]int{}, track...))
		return
	}
	for _, num := range nums {
		if !used[num] {
			track, used[num] = append(track, num), true
			backtrace(nums, track, used, ans)
			track, used[num] = track[:len(track)-1], false
		}
	}
}
