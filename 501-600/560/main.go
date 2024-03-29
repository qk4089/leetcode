package _60

//https://leetcode.cn/problems/subarray-sum-equals-k/
//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//子数组是数组中元素的连续非空序列。

//示例 1：
//输入：nums = [1,1,1], k = 2
//输出：2

//输入：nums = [1,2,3], k = 3
//输出：2

//提示：
//	1 <= nums.length <= 2 * 10^4
//	-1000 <= nums[i] <= 1000
//	-10^7 <= k <= 10^7

func subarraySum(nums []int, k int) int {
	count, preSum, hash := 0, 0, map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if hash[preSum-k] > 0 {
			count += hash[preSum-k]
		}
		hash[preSum]++
	}
	return count
}
