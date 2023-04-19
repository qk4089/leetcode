package _322

//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//你可以认为每种硬币的数量是无限的。

//示例
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1

//输入：coins = [2], amount = 3
//输出：-1

//输入：coins = [1], amount = 0
//输出：0

// 提示
//	1 <= coins.length <= 12
//	1 <= coins[i] <= 2^31 - 1
//	0 <= amount <= 10^4

// F(N) = min(F(N-1),F(N-2),F(N-5)) + 1
func coinChange(coins []int, amount int) int {
	result := make([]int, amount+1)
	for i := 1; i < len(result); i++ {
		result[i] = 10001
	}
	for i := 1; i < len(result); i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				result[i] = min(result[i], result[i-coins[j]]+1)
			}
		}
	}
	if result[amount] > 10000 {
		return -1
	}
	return result[amount]
}

func min(x, y int) int {
	if x != 0 && x < y {
		return x
	}
	return y
}
