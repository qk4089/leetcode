package beibao

import "math"

// 给定N个物品装入容量为W的背包，最大能装多满？

// 回溯

func recall(w int, items []int) int {
	// i表示考察到哪个物品了; cw表示当前已经装进去的物品的重量和; w背包重量; items表示每个物品的重量
	return doRecall(0, 0, w, items)
}

func doRecall(i, cw, w int, items []int) int {
	// 正常来说递归的终止条件是 i == len(items)，加上 cw == w 可以提前结束递归
	if cw == w || i == len(items) {
		return cw
	}
	maxV := doRecall(i+1, cw, w, items) // 不拿当前
	if cw+items[i] <= w {
		maxV = max(maxV, doRecall(i+1, cw+items[i], w, items)) // 拿当前（剪枝优化）
	}
	return maxV
}

// 动态规划

func dynamic(w int, items []int) int {
	result := make([][]bool, len(items))
	for i := 0; i < len(result); i++ {
		result[i] = make([]bool, w+1)
	}
	result[0][0] = true
	result[0][items[0]] = true
	for i := 1; i < len(items); i++ {
		for j := 0; j < w+1; j++ {
			if result[i-1][j] {
				result[i][j] = true
				if j+items[i] <= w {
					result[i][items[i]+j] = true
				}
			}
		}
	}
	for i := w; i >= 0; i-- {
		if result[len(items)-1][i] {
			return i
		}
	}
	return 0
}

// 当前层状态仅跟上层有关，可压缩空间
func dynamic2(w int, items []int) int {
	result := make([]bool, w+1)
	result[0] = true
	for i := 0; i < len(items); i++ {
		// 倒序遍历，不然会重复计算导致BUG
		for j := w - items[i]; j >= 0; j-- {
			if result[j] {
				result[j+items[i]] = true
			}
		}
	}
	for i := w; i >= 0; i-- {
		if result[i] {
			return i
		}
	}
	return 0
}

// 给定N个物品装入容量为W的背包，每个物品价值不等，在不超过背包容量的情况下最大价值？

// F(i, w) = max { F(i-1, w-cw) + vi，F(i-1, w) }
func dynamic_with_price(w int, items, values []int) int {
	result := make([][]int, len(items))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, w+1)
	}
	for i := 0; i < w+1; i++ {
		if i >= items[0] {
			result[0][i] = values[0]
		}
	}
	for i := 1; i < len(items); i++ {
		for j := 0; j < w+1; j++ {
			if j >= items[i] {
				result[i][j] = max(result[i-1][j-items[i]]+values[i], result[i-1][j])
			} else {
				result[i][j] = result[i-1][j]
			}
		}
	}
	return result[len(items)-1][w]
}

// 使用滚动数组优化
func dynamic_with_price2(m int, a []int) int {
	return 0
}

// 进一维压缩
func dynamic_with_price3(w int, items, values []int) int {
	result := make([]int, w+1)
	for i := 0; i < len(items); i++ {
		for j := w; j >= items[i]; j-- {
			result[j] = max(result[j-items[i]]+values[i], result[j])
		}
	}
	return result[w]
}

func max(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
