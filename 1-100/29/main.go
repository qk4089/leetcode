package _29

import "math"

//https://leetcode.cn/problems/divide-two-integers/description/
//给你两个整数，被除数dividend和除数divisor。将两数相除，要求不使用乘法、除法和取余运算。整数除法应该向零截断，也就是截去（truncate）其小数部分。
//例如，8.345将被截断为8，-2.7335将被截断至-2。返回被除数dividend除以除数divisor得到的商 。

//注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是[−2^31, 2^31−1]。本题中，如果商严格大于2^31−1，则返回2^31−1；
//如果商严格小于-2^31，则返回-2^31 。

//示例 1:
//输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = 3.33333.. ，向零截断后得到 3 。

//输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = -2.33333.. ，向零截断后得到 -2 。

//提示：
//	-2^31 <= dividend, divisor <= 2^31 - 1
//	divisor != 0

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 1 {
		return dividend
	}
	symbol, positive := dividend>>31 != divisor>>31, dividend>>31 == 0
	if symbol {
		divisor = 0 - divisor
	}
	ans, sum := 0, divisor
	for ans < math.MaxInt32 && (positive && dividend >= divisor || !positive && dividend <= divisor) {
		cnt := 1
		for positive && sum+sum <= dividend || !positive && sum+sum >= dividend {
			sum, cnt = sum+sum, cnt+cnt
		}
		ans, sum, dividend = ans+cnt, divisor, dividend-sum
	}
	if symbol {
		ans = 0 - ans
	}
	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	return ans
}
