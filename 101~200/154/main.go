package _154

import "sort"

//已知一个长度为n的数组，预先按照升序排列，经由1到n次旋转后，得到输入数组。
//例如，原数组nums=[0,1,4,4,5,6,7]在变化后可能得到：
//	若旋转4次，则可以得到[4,5,6,7,0,1,4]
//	若旋转7次，则可以得到[0,1,4,4,5,6,7]
//注意，数组[a[0],a[1],a[2],...,a[n-1]]旋转一次的结果为数组[a[n-1],a[0],a[1],a[2],...,a[n-2]]。
//给你一个可能存在重复元素值的数组nums，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的最小元素。
//你必须尽可能减少整个过程的操作步骤。

//进阶：这道题与寻找旋转排序数组中的最小值类似，但nums可能包含重复元素。允许重复会影响算法的时间复杂度吗？会如何影响，为什么？

//提示：
//	n==nums.length
//	1<=n<=5000
//	-5000<=nums[i]<=5000
//	nums原来是一个升序排序的数组，并进行了1至n次旋转

//示例
//输入：nums=[1,3,5]
//输出：1

//输入：nums=[2,2,2,0,1]
//输出：0

func findMin(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return nums[left]
}
