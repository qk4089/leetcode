package _12

//https://leetcode.cn/problems/integer-to-roman/
//罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
//	字符          数值
//	I             1
//	V             5
//	X             10
//	L             50
//	C             100
//	D             500
//	M             1000
//例如: 罗马数字2写做II，即为两个并列的1。12写做XII，即为X+II。27写做XXVII, 即为XX+V+II。通常情况下，罗马数字中小的数字在大的数字的右边。
//但也存在特例，例如4不写做IIII，而是IV。数字1在数字5的左边，所表示的数等于大数5减小数1得到的数值4。同样地，数字9表示为IX。
//这个特殊的规则只适用于以下六种情况：
//	I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//	X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//	C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

//给你一个整数，将其转为罗马数字。

//示例 1:
//输入: num = 3
//输出: "III"

//输入: num = 4
//输出: "IV"

//输入: num = 9
//输出: "IX"

//输入: num = 58
//输出: "LVIII"
//解释: L = 50, V = 5, III = 3.

//输入: num = 1994
//输出: "MCMXCIV"
//解释: M = 1000, CM = 900, XC = 90, IV = 4.

//提示：
//	1 <= num <= 3999

func intToRoman(num int) string {
	Q := []string{"", "M", "MM", "MMM"}
	B := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	S := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	G := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return Q[num/1000] + B[(num%1000)/100] + S[(num%100)/10] + G[num%10]
}
