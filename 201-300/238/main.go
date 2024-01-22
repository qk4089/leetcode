package _38

//https://leetcode.cn/problems/product-of-array-except-self/
//给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

//请不要使用除法，且在 O(n) 时间复杂度内完成此题。

//示例 1:
//输入: nums = [1,2,3,4]
//输出: [24,12,8,6]

//输入: nums = [-1,1,0,-3,3]
//输出: [0,0,9,0,0]

//提示：
//	2 <= nums.length <= 10^5
//	-30 <= nums[i] <= 30
//	保证数组nums之中任意元素的全部前缀元素和后缀的乘积都在32位整数范围内

//进阶：你可以在O(1)的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	revert(0, 1, nums, ans)
	return ans
}

func revert(idx, pre int, nums, ans []int) int {
	if idx+1 >= len(nums) {
		ans[idx] = pre
		return nums[idx]
	}
	suffix := revert(idx+1, pre*nums[idx], nums, ans)
	ans[idx] = pre * suffix
	return nums[idx] * suffix
}
