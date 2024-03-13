package _67

//给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

//示例 1：
//输入:a = "11", b = "1"
//输出："100"

//输入：a = "1010", b = "1011"
//输出："10101"

//提示：
//	1 <= a.length, b.length <= 104
//	a 和 b 仅由字符 '0' 或 '1' 组成
//	字符串如果不是 "0" ，就不含前导零

func addBinary(a string, b string) string {
	pa, pb, flag, ans := len(a)-1, len(b)-1, byte(0), ""
	for pa >= 0 || pb >= 0 || flag > 0 {
		var x byte = 0
		if pa >= 0 {
			x, pa = a[pa]-'0', pa-1
		}
		var y byte = 0
		if pb >= 0 {
			y, pb = b[pb]-'0', pb-1
		}
		flag, ans = (x+y+flag)/2, string(x^y^flag+'0')+ans
	}
	return ans
}
