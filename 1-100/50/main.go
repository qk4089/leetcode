package _50

//https://leetcode.cn/problems/powx-n/
//实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，x^n ）。

//示例 1：
//
//输入：x = 2.00000, n = 10
//输出：1024.00000

//输入：x = 2.10000, n = 3
//输出：9.26100

//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2^-2 = 1/2^2 = 1/4 = 0.25

//提示：
//	-100.0 < x < 100.0
//	-2^31 <= n <= 2^31-1
//	n 是一个整数
//	要么 x 不为零，要么 n > 0 。
//	-10^4 <= x^n <= 10^4

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	ans, cnt, positive := 1.0, 1, n>>31 == 0
	if !positive {
		n *= -1
	}
	for cnt <= n {
		tmp := x
		for 2*cnt <= n {
			cnt, tmp = 2*cnt, tmp*tmp
		}
		ans, cnt, n = ans*tmp, 1, n-cnt
	}
	if positive {
		return ans
	}
	return 1.0 / ans
}
