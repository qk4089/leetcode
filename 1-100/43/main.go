package _43

import "strconv"

//https://leetcode.cn/problems/multiply-strings/
//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。

//示例 1:
//输入: num1 = "2", num2 = "3"
//输出: "6"

//输入: num1 = "123", num2 = "456"
//输出: "56088"

//提示：
//	1 <= num1.length, num2.length <= 200
//	num1 和 num2 只能由数字组成。
//	num1 和 num2 都不包含任何前导零，除了数字0本身。

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := "0"
	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- {
		curr := ""
		add := 0
		for j := n - 1; j > i; j-- {
			curr += "0"
		}
		y := int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- {
			x := int(num1[j] - '0')
			product := x*y + add
			curr = strconv.Itoa(product%10) + curr
			add = product / 10
		}
		for ; add != 0; add /= 10 {
			curr = strconv.Itoa(add%10) + curr
		}
		ans = addStrings(ans, curr)
	}
	return ans
}

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
