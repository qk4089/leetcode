package _169

//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。

//示例 1：
//输入：nums = [3,2,3]
//输出：3

//输入：nums = [2,2,1,1,1,2,2]
//输出：2

//提示：
//	n == nums.length
//	1 <= n <= 5 * 10^4
//	-10^9 <= nums[i] <= 10^9

//进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

func majorityElement(nums []int) int {
	result, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			result, count = nums[i], 1
			continue
		}
		if nums[i] == result {
			count++
		} else {
			count--
		}
	}
	return result
}
