package _229

//给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

//示例
//输入：nums = [3,2,3]
//输出：[3]

//输入：nums = [1]
//输出：[1]

//输入：nums = [1,2]
//输出：[1,2]

//提示：
//	1 <= nums.length <= 5 * 10^4
//	-10^9 <= nums[i] <= 10^9

//进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。

func majorityElement(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	num1, vote1, num2, vote2 := 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if vote1 > 0 && nums[i] == num1 {
			vote1++
		} else if vote2 > 0 && nums[i] == num2 {
			vote2++
		} else if vote1 == 0 {
			num1, vote1 = nums[i], 1
		} else if vote2 == 0 {
			num2, vote2 = nums[i], 1
		} else {
			vote1, vote2 = vote1-1, vote2-1
		}
	}
	count1, count2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if vote1 > 0 && nums[i] == num1 {
			count1++
		} else if vote2 > 0 && nums[i] == num2 {
			count2++
		}
	}
	result := make([]int, 0)
	if count1 > len(nums)/3 {
		result = append(result, num1)
	}
	if count2 > len(nums)/3 {
		result = append(result, num2)
	}
	return result
}
