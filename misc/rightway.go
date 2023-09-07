package misc

import "fmt"

//N头牛排成一列1≤N≤5000。每头牛或者向前或者向后。为了让所有牛都面向前方，农夫每次可以将 K 头连续的牛转向，1≤K≤N，
//求使操作次数最小的相应 K 和最小的操作次数 M。F为朝前，B为朝后。

// 示例
// input: BBFBFBB
// return: 3, 3
// (1,2,3) (3,4,5) (5,6,7)

func flipCount(str string) (int, int) {
	minK, ans := 0, make([]int, len(str))
	for k := 0; k < len(str); k++ {
		ans[k] = flip(k+1, []byte(str))
		if ans[k] < ans[minK] {
			minK = k
		}
	}
	fmt.Println(ans)
	return minK + 1, ans[minK]
}

func flip(k int, str []byte) int {
	idx, count := 0, 0
	for ; idx < len(str)-k+1; idx++ {
		if str[idx] == 'F' {
			continue
		}
		for j := idx; j < idx+k; j++ {
			if str[j] == 'F' {
				str[j] = 'B'
			} else {
				str[j] = 'F'
			}
		}
		count++
	}
	for ; idx < len(str); idx++ {
		if str[idx] == 'B' {
			return 5001
		}
	}
	return count
}
