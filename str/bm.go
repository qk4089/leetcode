package str

// 坏字符、好后缀原则
func bm(str string, target string) int {
	idx, bc, suffix, prefix := 0, [256]int{}, make([]int, len(target)), make([]bool, len(target))
	for i := 0; i < 256; i++ {
		bc[i] = -1
	}
	for i, c := range target {
		bc[c] = i
	}
	generateGS(target, suffix, prefix)

	for idx <= len(str)-len(target) {
		j := len(target) - 1
		for ; j >= 0; j-- {
			if str[idx+j] != target[j] {
				break
			}
		}
		if j < 0 {
			return idx
		}
		idx += j - bc[str[idx+j]]
	}
	return -1
}

func generateGS(str string, suffix []int, prefix []bool) {
	m := len(str)
	for i := 0; i < m; i++ {
		suffix[i], prefix[i] = -1, false
	}

	for i := 0; i < m-1; i++ {
		j, k := i, 0
		for j >= 0 && str[j] == str[m-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}
		if j == -1 {
			prefix[k] = true
		}
	}
}
