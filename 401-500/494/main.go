package _494

//https://leetcode.cn/problems/target-sum/
//给你一个非负整数数组 nums 和一个整数 target 。向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

//示例 1：
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3

//输入：nums = [1], target = 1
//输出：1

//提示：
//	1 <= nums.length <= 20
//	0 <= nums[i] <= 1000
//	0 <= sum(nums[i]) <= 1000
//	-1000 <= target <= 1000

func findTargetSumWays(nums []int, target int) int {
	sums := 0
	for _, n := range nums {
		sums += n
	}
	if (sums+target)%2 > 0 || sums+target < 0 {
		return 0
	}
	avg, dp := (sums+target)/2, make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, avg+1)
	}
	dp[0][0] = 1
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= avg; j++ {
			if nums[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][avg]
}

func recall(nums []int, target int) int {
	cnt, sum := 0, 0
	var backtrace func(idx, sum int)
	backtrace = func(idx, sum int) {
		if idx == len(nums) {
			if sum == target {
				cnt++
			}
			return
		}
		backtrace(idx+1, sum+nums[idx])
		backtrace(idx+1, sum-nums[idx])
	}
	backtrace(0, sum)
	return cnt
}
