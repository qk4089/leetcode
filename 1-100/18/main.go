package _18

import "sort"

//https://leetcode.cn/problems/4sum/
//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组
//[nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//	0 <= a, b, c, d < n
//	a、b、c 和 d 互不相同
//	nums[a] + nums[b] + nums[c] + nums[d] == target
//你可以按 任意顺序 返回答案 。

//示例 1：
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]

//提示：
//	1 <= nums.length <= 200
//	-10^9 <= nums[i] <= 10^9
//	-10^9 <= target <= 10^9

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for _, val := range threeSum(nums[i+1:], target-nums[i]) {
			ans = append(ans, append([]int{nums[i]}, val...))
		}
	}
	return ans
}

func threeSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for _, val := range twoSum(nums[i+1:], target-nums[i]) {
			result = append(result, append([]int{nums[i]}, val...))
		}
	}
	return result
}

func twoSum(nums []int, target int) [][]int {
	left, right, result := 0, len(nums)-1, make([][]int, 0)
	for left < right {
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
