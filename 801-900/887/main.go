package _887

//给你k枚相同的鸡蛋，并可以使用一栋从第1层到第n层的建筑。已知存在楼层f，满足0 <= f <= n，任何从高于f的楼层落下的鸡蛋都会碎，
//从f楼层或比它低的楼层落下的鸡蛋都不会破。
//每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。
//如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。
//
//请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？

//示例 1：
//输入：k = 1, n = 2
//输出：2
//解释：
//鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。
//否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。
//如果它没碎，那么肯定能得出 f = 2 。
//因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。

//输入：k = 2, n = 6
//输出：3

//输入：k = 3, n = 14
//输出：4

//提示：
//	1 <= k <= 100
//	1 <= n <= 10^4

func superEggDrop(k int, n int) int {
	m, dp := 0, make([]int, k+1)
	for dp[k] < n {
		m++
		for i := k; i > 0; i-- {
			dp[i] = dp[i] + dp[i-1] + 1
		}
	}
	return m
}

//func dp(k int, n int, cache map[string]int) int {
//	if k == 1 {
//		return n
//	}
//	if n == 0 {
//		return 0
//	}
//	if val, ok := cache[strconv.Itoa(k)+","+strconv.Itoa(n)]; ok {
//		return val
//	}
//	ans := math.MaxInt32
//	for i := 1; i <= n; i++ {
//		ans = min(ans, max(dp(k-1, i-1, cache), dp(k, n-i, cache))+1)
//	}
//	cache[strconv.Itoa(k)+","+strconv.Itoa(n)] = ans
//	return ans
//}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
