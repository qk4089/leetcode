package _415

//https://leetcode.cn/problems/add-strings/
//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
//你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。

//示例 1：
//输入：num1 = "11", num2 = "123"
//输出："134"

//输入：num1 = "456", num2 = "77"
//输出："533"

//输入：num1 = "0", num2 = "0"
//输出："0"

//提示：
//	1 <= num1.length, num2.length <= 10^4
//	num1 和num2 都只包含数字 0-9
//	num1 和num2 都不包含任何前导零

func addStrings(num1 string, num2 string) string {
	var ZERO, flag byte = '0', 0
	p0, p1, ans := len(num1)-1, len(num2)-1, ""
	for p0 >= 0 || p1 >= 0 {
		x, y := ZERO, ZERO
		if p0 >= 0 {
			x, p0 = num1[p0], p0-1
		}
		if p1 >= 0 {
			y, p1 = num2[p1], p1-1
		}
		tmp := x + y - 2*ZERO + flag
		flag, ans = tmp/10, string(tmp%10+ZERO)+ans
	}
	if flag > 0 {
		ans = string('1') + ans
	}
	return ans
}
