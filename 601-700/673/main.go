package _673

//给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。
//注意 这个数列必须是 严格 递增的。
//
//示例 1:
//输入: [1,3,5,4,7]
//输出: 2
//解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
//
//输入: [2,2,2,2,2]
//输出: 5
//解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
//
//提示:
//	1 <= nums.length <= 2000
//	-10^6 <= nums[i] <= 10^6

func findNumberOfLIS(nums []int) int {
	maxLen, dp, cnt := 0, make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cnt[i], dp[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		if maxLen == dp[i] {
			ans += cnt[i]
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
