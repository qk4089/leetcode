package _8

import "math"

//请你来实现一个myAtoi(strings)函数，使其能将字符串转换成一个32位有符号整数（类似C/C++中的atoi函数）。
//函数myAtoi(strings)的算法如下：
//1.读入字符串并丢弃无用的前导空格
//2.检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。确定最终结果是负数还是正数。如果两者都不存在，则假定结果为正。
//3.读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
//4.将前面步骤读入的这些数字转换为整数（即，"123"->123，"0032"->32）。如果没有读入数字，则整数为0。必要时更改符号（从步骤2开始）。
//5.如果整数数超过32位有符号整数范围[−2^31,2^31−1]，需要截断这个整数，使其保持在这个范围内。具体来说，小于−2^31的整数应该被固定为−2^31，大于2^31−1的整数应该被固定为2^31−1。
//返回整数作为最终结果。
//注意：
//	本题中的空白字符只包括空格字符' '。
//	除前导空格或数字后的其余字符串外，请勿忽略任何其他字符。

//提示
//	0 <= s.length <= 200
//	s 由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成

//示例 1：
//输入：s = "42"
//输出：42
//解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
//第 1 步："42"（当前没有读入字符，因为没有前导空格）
//第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//第 3 步："42"（读入 "42"）
//解析得到整数 42, 由于 "42" 在范围 [-231, 231 - 1] 内，最终结果为 42 。

//输入：s = "   -42"
//输出：-42
//解释：
//第 1 步："   -42"（读入前导空格，但忽视掉）
//第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
//第 3 步："   -42"（读入 "42"）
//解析得到整数 -42 。
//由于 "-42" 在范围 [-231, 231 - 1] 内，最终结果为 -42 。

//输入：s = "4193 with words"
//输出：4193
//解释：
//第 1 步："4193 with words"（当前没有读入字符，因为没有前导空格）
//第 2 步："4193 with words"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//第 3 步："4193 with words"（读入 "4193"；由于下一个字符不是一个数字，所以读入停止）
//解析得到整数 4193 。
//由于 "4193" 在范围 [-231, 231 - 1] 内，最终结果为 4193 。

func myAtoi(s string) int {
	point, sb := 0, []byte(s)
	for ; point < len(sb); point++ {
		if sb[point] != ' ' {
			break
		}
	}
	if point == len(sb) {
		return 0
	}
	start, flag := 0, 1
	if sb[point] == '-' {
		flag, point = -1, point+1
	} else if sb[point] == '+' {
		point++
	}
	for ; point < len(sb); start, point = start+1, point+1 {
		if sb[point] < '0' || sb[point] > '9' {
			break
		}
		sb[start] = sb[point] - '0'
	}
	if start == 0 {
		return 0
	}
	sum, power, overflow := 0, func(n int) int { return int(math.Pow(float64(10), float64(n))) }, func() int {
		if flag == 1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}()
	for i, b := range sb[:start] {
		base := power(start - i - 1)
		if (base < 0 || base > math.MaxInt32) && int(b) != 0 {
			return overflow
		}
		tmp := int(b) * base
		if tmp > math.MaxInt32 {
			return overflow
		}
		sum += tmp
	}
	if sum > math.MaxInt32 {
		return overflow
	}
	return sum * flag
}
