package _152

//https://leetcode.cn/problems/maximum-product-subarray/
//给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//测试用例的答案是一个 32-位 整数。

//子数组 是数组的连续子序列。

//示例 1:
//输入: nums = [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。

//输入: nums = [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

//提示:
//	1 <= nums.length <= 2 * 10^4
//	-10 <= nums[i] <= 10
//	nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数

func maxProduct(nums []int) int {
	maxV, maxTmp, minTmp := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxTmp, minTmp = max(max(maxTmp*nums[i], minTmp*nums[i]), nums[i]), min(min(maxTmp*nums[i], minTmp*nums[i]), nums[i])
		maxV = max(maxV, maxTmp)
	}
	return maxV
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
