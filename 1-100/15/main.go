package _15

import "sort"

//https://leetcode.cn/problems/3sum/
//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
//同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//注意：答案中不可以包含重复的三元组。

//示例
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。

//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。

//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。

//提示：
//	3 <= nums.length <= 3000
//	-10^5 <= nums[i] <= 10^5

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for _, val := range twoSum(nums[i+1:], 0-nums[i]) {
			ans = append(ans, append([]int{nums[i]}, val...))
		}
	}
	return ans
}

func twoSum(nums []int, target int) [][]int {
	left, right, result := 0, len(nums)-1, make([][]int, 0)
	for nums[left] <= target && left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			result = append(result, []int{nums[left], nums[right]})
			left, right = left+1, right-1
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return result
}
