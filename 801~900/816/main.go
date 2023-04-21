package _816

import "bytes"

//我们有一些二维坐标，如"(1, 3)"或"(2, 0.5)"，然后我们移除所有逗号，小数点和空格，得到一个字符串S，返回所有可能的原始字符串到一个列表中

//原始的坐标表示法不会存在多余的零，所以不会出现类似于"00", "0.0", "0.00", "1.0", "001", "00.01"或一些其他更小的数来表示坐标。
//此外，一个小数点前至少存在一个数，所以也不会出现“.1”形式的数字。最后返回的列表可以是任意顺序的。而且注意返回的两个数字中间（逗号之后）都有一个空格
//提示:
//	4 <= S.length <= 12
//	S[0] = "(", S[S.length - 1] = ")", 且字符串 S 中的其他元素都是数字

//示例
//输入: "(123)"
//输出: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]

//输入: "(00011)"
//输出:  ["(0.001, 1)", "(0, 0.011)"]
//解释: 0.0, 00, 0001 或 00.01 是不被允许的

//输入: "(0123)"
//输出: ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"]

//输入: "(100)"
//输出: [(10, 0)]
//解释:1.0 是不被允许的

func ambiguousCoordinates(s string) []string {
	b, result := []byte(s)[1:len(s)-1], make([]string, 0)
	l := len(b)
	var buffer bytes.Buffer
	for i := 1; i < l; i++ {
		left, right := b[0:i], b[i:l]
		for _, v1 := range generator(&buffer, left) {
			for _, v2 := range generator(&buffer, right) {
				result = append(result, concat(&buffer, v1, v2))
			}
		}
	}
	return result
}

func generator(buffer *bytes.Buffer, b []byte) []string {
	if len(b) == 1 {
		return []string{string(b)}
	}
	var result []string
	if b[0] != '0' {
		result = append(result, string(b))
		if b[len(b)-1] != '0' {
			for i := len(b) - 1; i > 0; i-- {
				buffer.Write(b[:i])
				buffer.WriteByte('.')
				buffer.Write(b[i:])
				result = append(result, buffer.String())
				buffer.Reset()
			}
		}
	} else if b[0] == '0' && b[len(b)-1] != '0' {
		buffer.WriteByte('0')
		buffer.WriteByte('.')
		buffer.Write(b[1:])
		result = append(result, buffer.String())
		buffer.Reset()
	}
	return result
}

func concat(buffer *bytes.Buffer, s1 string, s2 string) string {
	defer buffer.Reset()
	buffer.WriteByte('(')
	buffer.WriteString(s1)
	buffer.WriteByte(',')
	buffer.WriteByte(' ')
	buffer.WriteString(s2)
	buffer.WriteByte(')')
	return buffer.String()
}
