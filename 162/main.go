package _162

//峰值元素是指其值严格大于左右相邻值的元素。
//给你一个整数数组nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
//你可以假设nums[-1]=nums[n]=-∞。你必须实现时间复杂度为O(logn)的算法来解决此问题。

//提示：
//	1<=nums.length<=1000
//	-2^31<=nums[i]<=2^31-1
//	对于所有有效的i都有nums[i]!=nums[i+1]

//示例
//输入：nums=[1,2,3,1]
//输出：2
//解释：3是峰值元素，你的函数应该返回其索引2。

//输入：nums=[1,2,1,3,5,6,4]
//输出：1或5
//解释：你的函数可以返回索引1，其峰值元素为2；或者返回索引5，其峰值元素为6。

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for right-left > 1 {
		mid := left + (right-left)>>1
		if nums[mid] < nums[mid-1] {
			right = mid - 1
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	if nums[left] < nums[right] {
		return right
	} else {
		return left
	}
}
