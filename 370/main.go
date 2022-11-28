package _370

//假设你有一个长度为n的数组，初始情况下所有数组都是0，你将会被给出k个更新动作。
//其中每个操作用一组三元组表示：[start, end, inc]，你需要将数组中元素（包括start和end）增加inc，请返回k次操作后的数组

//示例
//输入：length=5,updates=[[1, 3, 2], [2, 4, 3], [0, 2, -2]]
//输出：[-2, 0, 3, 5, 3]
//解释：
//	初始状态：[0, 0, 0, 0, 0]
//	执行操作[1, 3, 2]后：[0, 2, 2, 2, 0]
//	执行操作[2, 4, 3]后：[0, 2, 5, 5, 3]
//	执行操作[0, 2, -2]后：[-2, 0, 3, 5, 3]

func getModifiedArray(length int, updates [][]int) []int {
	diff := make([]int, length)
	for _, update := range updates {
		left, right, val := update[0], update[1], update[2]
		diff[left] += val
		if right+1 < length {
			diff[right+1] -= val
		}
	}
	return restore(diff)
}

func restore(diff []int) []int {
	for i := 1; i < len(diff); i++ {
		diff[i] = diff[i-1] + diff[i]
	}
	return diff
}
