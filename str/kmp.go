package str

func kmp(origin string, target string) int {
	point, ro, rt := 0, []rune(origin), []rune(target)
	next := getNext(rt)
	for i := 0; i < len(ro); i++ {
		// 尽可能把 point 移到正确的位置（最差从0开始）继续比较
		for point > 0 && ro[i] != rt[point] {
			point = next[point-1] + 1
		}
		if ro[i] == rt[point] {
			point++
		}
		if point == len(rt) {
			return i - len(rt) + 1
		}
	}
	return -1
}

func getNext(str []rune) []int {
	k, result := -1, make([]int, len(str))
	result[0] = -1
	for i := 1; i < len(str); i++ {
		for k != -1 && str[k+1] != str[i] {
			k = result[k]
		}
		if str[k+1] == str[i] {
			k++
		}
		result[i] = k
	}
	return result
}
