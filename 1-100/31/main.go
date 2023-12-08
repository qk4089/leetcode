package _31

import "sort"

//https://leetcode.cn/problems/next-permutation/
//整数数组的一个排列就是将其所有成员以序列或线性顺序排列。例如arr=[1,2,3]，以下这些都可以视作arr的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1]。
//整数数组的 下一个排列是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，
//那么数组的 下一个排列就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

//例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
//类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
//而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。

//给你一个整数数组 nums ，找出 nums 的下一个排列。必须 原地 修改，只允许使用额外常数空间。

//示例 1：
//输入：nums = [1,2,3]
//输出：[1,3,2]

//输入：nums = [3,2,1]
//输出：[1,2,3]

//输入：nums = [1,1,5]
//输出：[1,5,1]

//提示：
//	1 <= nums.length <= 100
//	0 <= nums[i] <= 100

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	minIdx := len(nums) - 1
	for idx := len(nums) - 1; idx > 0; idx-- {
		if nums[idx] <= nums[idx-1] {
			continue
		}
		for nums[minIdx] <= nums[idx-1] && minIdx > idx-1 {
			minIdx--
		}
		nums[minIdx], nums[idx-1] = nums[idx-1], nums[minIdx]
		sortInPlace(nums[idx:])
		return
	}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func sortInPlace(nums []int) {
	sort.Ints(nums)
}
