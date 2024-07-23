package _174

//https://leetcode.cn/problems/dungeon-game/
//恶魔们抓住了公主并将她关在了地下城dungeon的右下角。地下城是由 m x n 个房间组成的二维网格。我们英勇的骑士最初被安置在左上角的房间里， 他必
//须穿过地下城并通过对抗恶魔来拯救公主。骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。 有些房间由恶魔
//守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值 为 0），要么
//包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。为了尽快解救公主，骑士决定 每次只 向右 或 向下 移动一步。
//返回确保骑士能够拯救到公主所需的最低初始健康点数。

//注意：任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。

//示例
//输入：dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
//输出：7
//解释：如果骑士遵循最佳路径：右 -> 右 -> 下 -> 下 ，则骑士的初始健康点数至少为 7 。

//输入：dungeon = [[0]]
//输出：1

//提示：
//	m == dungeon.length
//	n == dungeon[i].length
//	1 <= m, n <= 200
//	-1000 <= dungeon[i][j] <= 1000

func calculateMinimumHP(dungeon [][]int) int {
	m, n, dp := len(dungeon), len(dungeon[0]), make([][]int, len(dungeon))
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = generate(1, dungeon[m-1][n-1])
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = generate(dp[i+1][n-1], dungeon[i][n-1])
	}
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = generate(dp[m-1][i+1], dungeon[m-1][i])
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = min(generate(dp[i+1][j], dungeon[i][j]), generate(dp[i][j+1], dungeon[i][j]))
		}
	}
	return dp[0][0]
}

func generate(target, current int) int {
	val := target - current
	if val > 1 {
		return val
	}
	return 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
