package _41

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

//示例
//输入：nums = [1,2,0]
//输出：3

//输入：nums = [3,4,-1,1]
//输出：2

//输入：nums = [7,8,9,11,12]
//输出：1

//提示：
//	1 <= nums.length <= 5 * 10^5
//	-2^31 <= nums[i] <= 2^31 - 1

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		idx := (nums[i] - 1) % n
		for nums[i] > 0 && idx != i && ((nums[idx]-1)%n != idx || nums[i] < nums[idx]) {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx = (nums[i] - 1) % n
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}
	return n + 1
}
