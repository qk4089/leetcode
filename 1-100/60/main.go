package _60

//https://leetcode.cn/problems/permutation-sequence/
//给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
//按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
//	"123"
//	"132"
//	"213"
//	"231"
//	"312"
//	"321"
//给定 n 和 k，返回第 k 个排列。

//示例 1：
//输入：n = 3, k = 3
//输出："213"

//输入：n = 4, k = 9
//输出："2314"

//输入：n = 3, k = 1
//输出："123"

//提示：
//	1 <= n <= 9
//	1 <= k <= n!

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	list, ans := make([]byte, n), make([]byte, 0)
	for i := 0; i < n; i++ {
		list[i] = byte(i+1) + '0'
	}
	for k-1 > 0 {
		mod := computed(n - 1)
		quotient := (k - 1) / mod
		ans = append(ans, list[quotient])
		list = append(list[0:quotient], list[quotient+1:]...)
		k = ((k - 1) % mod) + 1
		n = n - 1
	}
	ans = append(ans, list...)
	return string(ans)
}

func computed(n int) int {
	if n == 1 {
		return 1
	}
	product := 1
	for i := 2; i <= n; i++ {
		product *= i
	}
	return product
}
