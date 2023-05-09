package _300

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

//示例 1：
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

//输入：nums = [0,1,0,3,2,3]
//输出：4

//输入：nums = [7,7,7,7,7,7,7]
//输出：1

//提示：
//	1 <= nums.length <= 2500
//	-10^4 <= nums[i] <= 10^4

//进阶：你能将算法的时间复杂度降低到 O(n log(n)) 吗?

func lengthOfLIS(nums []int) int {
	maxV, result := 1, make([]int, len(nums))
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				result[i] = max(result[i], result[j]+1)
			}
		}
		maxV = max(maxV, result[i])
	}
	return maxV
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
