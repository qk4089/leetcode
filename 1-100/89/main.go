package _89

import (
	"math"
)

//https://leetcode.cn/problems/gray-code/
//n 位格雷码序列 是一个由 2n 个整数组成的序列，其中: 每个整数都在范围 [0, 2n - 1] 内（含 0 和 2n - 1）
//第一个整数是 0, 一个整数在序列中出现 不超过一次
//每对相邻整数的二进制表示恰好一位不同 ，且第一个和最后一个整数的二进制表示恰好一位不同
//给你一个整数 n ，返回任一有效的 n 位格雷码序列 。

//示例 1：
//输入：n = 2
//输出：[0,1,3,2]
//解释：
//[0,1,3,2] 的二进制表示是 [00,01,11,10] 。
//- 00 和 01 有一位不同
//- 01 和 11 有一位不同
//- 11 和 10 有一位不同
//- 10 和 00 有一位不同
//[0,2,3,1] 也是一个有效的格雷码序列，其二进制表示是 [00,10,11,01] 。
//- 00 和 10 有一位不同
//- 10 和 11 有一位不同
//- 11 和 01 有一位不同
//- 01 和 00 有一位不同

//输入：n = 1
//输出：[0,1]

//提示：
//	1 <= n <= 16

func grayCode(n int) []int {
	size := pow(2, n)
	ans, used := make([]int, size), make([]bool, size)
	ans[0], used[0] = 0, true
	backtrace(1, used, ans)
	return ans
}

func backtrace(idx int, used []bool, ans []int) bool {
	if idx == len(ans) {
		return true
	}
	for i := 1; i < len(used); i++ {
		if !used[i] && isValid(i, idx, ans) {
			ans[idx], used[i] = i, true
			if backtrace(idx+1, used, ans) {
				return true
			}
			ans[idx], used[i] = 0, false
		}
	}
	return false
}

func isValid(val, idx int, ans []int) bool {
	n := len(ans)
	valid := func(x, y int) bool {
		v, tmp := 1, x^y
		for i := 0; i < n && v <= tmp; i++ {
			if v == tmp {
				return true
			}
			v = v << 1
		}
		return false
	}
	flag := valid(ans[idx-1], val)
	if flag && idx == len(ans)-1 {
		flag = valid(0, val)
	}
	return flag
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
