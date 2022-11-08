package _1678

//请你设计一个可以解释字符串command的Goal解析器。

//command由"G"、"()"和/或"(al)"按某种顺序组成。Goal解析器会将"G"解释为字符串"G"、"()"解释为字符串"o"，"(al)"解释为字符串"al"。
//然后按原顺序将经解释得到的字符串连接成一个字符串。给你字符串command，返回Goal解析器对command的解释结果。
//提示：
//	1 <= command.length <= 100
//	command 由 "G"、"()" 和/或 "(al)" 按某种顺序组成

//示例
//输入：command = "G()(al)"
//输出："Goal"

//输入：command = "G()()()()(al)"
//输出："Gooooal"

//输入：command = "(al)G(al)()()G"
//输出："alGalooG"

func main(command string) string {
	b := []byte(command)
	str := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == 'G' {
			str = append(str, 'G')
		} else if b[i] == '(' {
			i++
			if b[i] == ')' {
				str = append(str, 'o')
			} else {
				str = append(str, 'a', 'l')
				i = i + 2
			}
		}
	}
	return string(str)
}
