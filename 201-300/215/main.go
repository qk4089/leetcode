package _215

import "math/rand"

//https://leetcode.cn/problems/kth-largest-element-in-an-array/
//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。请注意，你需要找的是数组排序后的第k个最大的元素，而不是第 k 个不同的元素。

//你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

//示例 1:
//输入: [3,2,1,5,6,4], k = 2
//输出: 5

//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4

//提示：
//	1 <= k <= nums.length <= 10^5
//	-10^4 <= nums[i] <= 10^4

func findKthLargest(nums []int, k int) int {
	return kSort(nums, len(nums)-k)
}

func kSort(nums []int, k int) int {
	idx := rand.Intn(len(nums))
	nums[idx], nums[len(nums)-1] = nums[len(nums)-1], nums[idx]
	point, start, pivot := 0, 0, nums[len(nums)-1]
	for point < len(nums)-1 {
		if nums[point] < pivot {
			nums[point], nums[start] = nums[start], nums[point]
			start++
		}
		point++
	}
	nums[point], nums[start] = nums[start], nums[point]
	if k == start {
		return nums[start]
	} else if k > start {
		for start < k && nums[start] == nums[start+1] {
			start++
		}
		if k == start {
			return nums[start]
		}
		return kSort(nums[start+1:], k-start-1)
	} else {
		return kSort(nums[:start], k)
	}
}
