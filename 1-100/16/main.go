package _16

import (
	"math"
	"sort"
)

//https://leetcode.cn/problems/3sum-closest/
//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。返回这三个数的和。

//假定每组输入只存在恰好一个解。

//示例 1：
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

//输入：nums = [0,0,0], target = 1
//输出：0

//提示：
//	3 <= nums.length <= 1000
//	-1000 <= nums[i] <= 1000
//	-10^4 <= target <= 10^4

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans, diff := 0, math.MaxInt
	for i := 0; i <= len(nums)-3; i++ {
		left, right, expect := i+1, len(nums)-1, target-nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			if sum == expect {
				return target
			} else if sum < expect {
				left++
			} else {
				right--
			}
			if abs(expect-sum) < diff {
				ans, diff = nums[i]+sum, abs(expect-sum)
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
