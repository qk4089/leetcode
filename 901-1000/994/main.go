package _994

//https://leetcode.cn/problems/rotting-oranges/
//在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：值 0 代表空单元格； 值 1 代表新鲜橘子； 值 2 代表腐烂的橘子。
//每分钟，腐烂的橘子周围4个方向上相邻的新鲜橘子都会腐烂。返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回-1。

//示例 1：
//输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
//输出：4

//输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
//输出：-1
//解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。

//输入：grid = [[0,2]]
//输出：0
//解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。

//提示：
//	m == grid.length
//	n == grid[i].length
//	1 <= m, n <= 10
//	grid[i][j] 仅为 0、1 或 2

func orangesRotting(grid [][]int) int {
	m, n, sum, queue := len(grid), len(grid[0]), 0, make([][2]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
			sum++
		}
	}
	if len(queue) == sum {
		return 0
	}
	start, cnt := 0, 0
	for len(queue) != sum && len(queue)-start > 0 {
		size := len(queue) - start
		for i := start; i < start+size; i++ {
			x, y := queue[i][0], queue[i][1]
			if x-1 >= 0 && grid[x-1][y] == 1 {
				queue, grid[x-1][y] = append(queue, [2]int{x - 1, y}), 2
			}
			if x+1 < m && grid[x+1][y] == 1 {
				queue, grid[x+1][y] = append(queue, [2]int{x + 1, y}), 2
			}
			if y-1 >= 0 && grid[x][y-1] == 1 {
				queue, grid[x][y-1] = append(queue, [2]int{x, y - 1}), 2
			}
			if y+1 < n && grid[x][y+1] == 1 {
				queue, grid[x][y+1] = append(queue, [2]int{x, y + 1}), 2
			}
		}
		cnt, start = cnt+1, start+size
	}
	if len(queue) == sum {
		return cnt
	} else {
		return -1
	}
}
