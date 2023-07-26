package _870

import "sort"

//给定两个长度相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的数目来描述。
//返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。

//示例 1：
//输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
//输出：[2,11,7,15]

//输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
//输出：[24,32,8,12]

//提示：
//	1 <= nums1.length <= 10^5
//	nums2.length == nums1.length
//	0 <= nums1[i], nums2[i] <= 10^9

func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	result := make([]int, len(nums1))
	for i := 0; i < len(nums2); i++ {
		idx := index(nums1, nums2[i])
		result[i] = nums1[idx]
		nums1 = append(nums1[0:idx], nums1[idx+1:]...)
	}
	return result
}

func index(nums []int, target int) int {
	if target < nums[0] || target >= nums[len(nums)-1] {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
