package _4

//https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//算法的时间复杂度应该为 O(log (m+n)) 。

//示例 1：
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2

//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

//提示：
//	nums1.length == m
//	nums2.length == n
//	0 <= m <= 1000
//	0 <= n <= 1000
//	1 <= m + n <= 2000
//	-10^6 <= nums1[i], nums2[i] <= 10^6

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	cnt := len(nums1) + len(nums2)
	mid, mod := cnt/2, cnt%2
	if mod == 0 {
		return (findKth(mid, nums1, nums2) + findKth(mid+1, nums1, nums2)) / 2.0
	} else {
		return findKth(mid+1, nums1, nums2)
	}
}

func findKth(k int, nums1, nums2 []int) float64 {
	for k > 1 && len(nums1) > 0 && len(nums2) > 0 {
		midK := min(min(len(nums1), len(nums2)), k/2)
		if nums1[midK-1] < nums2[midK-1] {
			nums1 = nums1[midK:]
		} else {
			nums2 = nums2[midK:]
		}
		k -= midK
	}
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	} else if len(nums2) == 0 {
		return float64(nums1[k-1])
	}
	return float64(min(nums1[0], nums2[0]))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
