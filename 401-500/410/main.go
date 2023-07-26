package _410

//给定一个非负整数数组 nums 和一个整数 m ，你需要将这个数组分成 m 个非空的连续子数组。
//设计一个算法使得这 m 个子数组各自和的最大值最小。

//示例 1：
//输入：nums = [7,2,5,10,8], m = 2
//输出：18
//解释：
//一共有四种方法将 nums 分割为 2 个子数组。
//其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
//因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。

//输入：nums = [1,2,3,4,5], m = 2
//输出：9

//输入：nums = [1,4,4], m = 3
//输出：4

//提示：
//	1 <= nums.length <= 1000
//	0 <= nums[i] <= 10^6
//	1 <= m <= min(50, nums.length)

func splitArray(nums []int, k int) int {
	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > left {
			left = nums[i]
		}
		right += nums[i]
	}
	for left < right {
		mid := left + (right-left)>>1
		tmp := count(nums, mid)
		if tmp > k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func count(nums []int, target int) int {
	sum, result := 0, 0
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] <= target {
			sum += nums[i]
		} else {
			result++
			sum = nums[i]
		}
	}
	return result + 1
}
