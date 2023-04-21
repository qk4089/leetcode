package _7

import (
	"math"
)

//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

//如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31−1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。

//示例
//输入：x = 123
//输出：321

//输入：x = -123
//输出：-321

//输入：x = 120
//输出：21

//输入：x = 0
//输出：0

//提示：
//	-2^31 <= x <= 2^31 - 1

func reverse(x int) int {
	if x > -10 && x < 10 {
		return x
	}
	flag := 1
	if x < 0 {
		flag, x = -1, x*-1
	}
	arr := make([]byte, 0)
	for x > 0 {
		x, arr = x/10, append(arr, byte(x%10))
	}
	sum := 0
	for _, b := range arr {
		sum = sum*10 + int(b)
		if sum > math.MaxInt32 {
			return 0
		}
	}
	return sum * flag
}

func sort() {

}
