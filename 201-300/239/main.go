package _239

import "math"

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
//你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。 返回滑动窗口中的最大值 。

//示例 1：
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7

//输入：nums = [1], k = 1
//输出：[1]

//提示：
//	1 <= nums.length <= 10^5
//	-10^4 <= nums[i] <= 10^4
//	1 <= k <= nums.length

func maxSlidingWindow(nums []int, k int) []int {
	result, maxV := make([]int, 0), math.MinInt
	for i := 0; i < k; i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
	}
	result = append(result, maxV)
	for i := k; i < len(nums); i++ {
		if nums[i] >= maxV {
			maxV = nums[i]
		} else if nums[i-k] >= maxV {
			maxV = math.MinInt
			for j := i - k + 1; j <= i; j++ {
				if nums[j] > maxV {
					maxV = nums[j]
				}
			}
		}
		result = append(result, maxV)
	}
	return result
}
