package _518

//https://leetcode.cn/problems/coin-change-ii/
//给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
//请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//假设每一种面额的硬币有无限个。题目数据保证结果符合 32 位带符号整数。

//示例 1：
//输入：amount = 5, coins = [1, 2, 5]
//输出：4
//解释：有四种方式可以凑成总金额：
//5=5
//5=2+2+1
//5=2+1+1+1
//5=1+1+1+1+1

//输入：amount = 3, coins = [2]
//输出：0
//解释：只用面额 2 的硬币不能凑成总金额 3 。

//输入：amount = 10, coins = [10]
//输出：1

//提示：
//	1 <= coins.length <= 300
//	1 <= coins[i] <= 5000
//	coins 中的所有值 互不相同
//	0 <= amount <= 5000

func change(amount int, coins []int) int {
	dp := make([][]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, len(coins)+1)
	}
	for i := 0; i <= len(coins); i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= amount; i++ {
		for j := 1; j <= len(coins); j++ {
			if coins[j-1] > i {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = dp[i-coins[j-1]][j] + dp[i][j-1]
			}
		}
	}
	return dp[amount][len(coins)]
}
