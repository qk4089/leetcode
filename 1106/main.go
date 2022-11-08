package _1106

//给你一个以字符串形式表述的布尔表达式（boolean）expression，返回该式的运算结果。

//有效的表达式需遵循以下约定：
//	"t"，运算结果为 True
//	"f"，运算结果为 False
//	"!(expr)"，运算过程为对内部表达式 expr 进行逻辑 非的运算（NOT）
//	"&(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 与的运算（AND）
//	"|(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 或的运算（OR）
//提示：
//	1 <= expression.length <= 20000
//	expression[i] 由 {'(', ')', '&', '|', '!', 't', 'f', ','} 中的字符组成。
//	expression 是以上述形式给出的有效表达式，表示一个布尔值。

//示例
//输入：expression = "!(f)"
//输出：true

//输入：expression = "|(f,t)"
//输出：true

//输入：expression = "&(t,f)"
//输出：false

//输入：expression = "|(&(t,f,t),!(t))"
//输出：false

func main(expression string) bool {
	b, stack := []byte(expression), &Stack{}
	for index := 0; index < len(b); index++ {
		char := b[index]
		if char == ',' {
			continue
		}
		if char == ')' {
			reduce(stack)
			continue
		}
		if char == '!' || char == '&' || char == '|' {
			index++
		}
		stack.Push(char)
	}
	for stack.Size() > 1 {
		reduce(stack)
	}
	return stack.Pop() == 't'
}

func reduce(stack *Stack) {
	var expr []byte
	element := stack.Pop()
	for element != '!' && element != '&' && element != '|' {
		expr = append(expr, element)
		element = stack.Pop()
	}
	switch element {
	case '!':
		stack.Push(not(expr))
	case '&':
		stack.Push(allOf(expr))
	case '|':
		stack.Push(anyOf(expr))
	}
}

func not(b []byte) byte {
	if b[0] == 'f' {
		return 't'
	} else {
		return 'f'
	}
}

func allOf(b []byte) byte {
	for _, v := range b {
		if v == 'f' {
			return 'f'
		}
	}
	return 't'
}

func anyOf(b []byte) byte {
	for _, v := range b {
		if v == 't' {
			return 't'
		}
	}
	return 'f'
}

type Stack struct {
	data []byte
}

func (s *Stack) Push(value byte) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() byte {
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *Stack) Size() int {
	return len(s.data)
}
