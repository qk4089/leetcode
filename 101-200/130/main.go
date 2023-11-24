package _130

//https://leetcode.cn/problems/surrounded-regions/
//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

//示例 1：
//输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
//输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
//最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

//输入：board = [["X"]]
//输出：[["X"]]

//提示：
//	m == board.length
//	n == board[i].length
//	1 <= m, n <= 200
//	board[i][j] 为 'X' 或 'O'

func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	visited := make([][]byte, n)
	for i := 0; i < len(board); i++ {
		visited[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		if board[0][i] == 'O' && visited[0][i] == 0 {
			visited[0][i] = 'O'
			bfs(0, i, visited, board)
		}
		if board[n-1][i] == 'O' && visited[n-1][i] == 0 {
			visited[n-1][i] = 'O'
			bfs(n-1, i, visited, board)
		}
	}
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' && visited[i][0] == 0 {
			visited[i][0] = 'O'
			bfs(i, 0, visited, board)
		}
		if board[i][m-1] == 'O' && visited[i][m-1] == 0 {
			visited[i][m-1] = 'O'
			bfs(i, m-1, visited, board)
		}
	}
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if visited[x][y] == 0 {
				board[x][y] = 'X'
			} else {
				board[x][y] = visited[x][y]
			}
		}
	}
}

func bfs(x, y int, visited [][]byte, board [][]byte) {
	queue := [][]int{{x, y}}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			tx, ty := queue[i][0], queue[i][1]
			// 上
			if tx >= 1 && visited[tx-1][ty] == 0 {
				visited[tx-1][ty] = board[tx-1][ty]
				if visited[tx-1][ty] == 'O' {
					queue = append(queue, []int{tx - 1, ty})
				}
			}
			// 下
			if tx+1 < len(board) && visited[tx+1][ty] == 0 {
				visited[tx+1][ty] = board[tx+1][ty]
				if visited[tx+1][ty] == 'O' {
					queue = append(queue, []int{tx + 1, ty})
				}
			}
			// 左
			if ty >= 1 && visited[tx][ty-1] == 0 {
				visited[tx][ty-1] = board[tx][ty-1]
				if visited[tx][ty-1] == 'O' {
					queue = append(queue, []int{tx, ty - 1})
				}
			}
			// 右
			if ty+1 < len(board[0]) && visited[tx][ty+1] == 0 {
				visited[tx][ty+1] = board[tx][ty+1]
				if visited[tx][ty+1] == 'O' {
					queue = append(queue, []int{tx, ty + 1})
				}
			}

		}
		queue = queue[size:]
	}
}
