package _34

//给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//如果数组中不存在目标值 target，返回 [-1, -1]。
//你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

//示例 1：
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]

//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]

//输入：nums = [], target = 0
//输出：[-1,-1]

//提示：
//	0 <= nums.length <= 10^5
//	-10^9 <= nums[i] <= 10^9
//	nums 是一个非递减数组
//	-10^9 <= target <= 10^9

func searchRange(nums []int, target int) []int {
	result := make([]int, 0)
	point := search(nums, target)
	if point == -1 {
		return []int{-1, -1}
	}
	ll, lr := 0, point
	for ll < lr {
		mid := ll + (lr-ll)>>1
		if nums[mid] == target {
			if mid == lr {
				break
			}
			lr = mid
		} else if nums[mid] < target {
			ll = mid + 1
		}
	}
	rl, rr := point, len(nums)-1
	for rl < rr {
		mid := rl + (rr-rl)>>1
		if nums[mid] == target {
			rl = mid + 1
		} else if nums[mid] > target {
			rr = mid - 1
		}
	}
	if nums[rl] != target {
		rl--
	}
	return append(result, ll, rl)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
