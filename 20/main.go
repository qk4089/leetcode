package _20

//给定一个只包括'('，')'，'{'，'}'，'['，']'的字符串s，判断字符串是否有效。
//有效字符串需满足：
//	左括号必须用相同类型的右括号闭合。
//	左括号必须以正确的顺序闭合。
//	每个右括号都有一个对应的相同类型的左括号。

//提示：
//	1<=s.length<=104
//	s仅由括号'()[]{}'组成

//示例
//输入：s = "()"
//输出：true
//
//输入：s = "()[]{}"
//输出：true
//
//输入：s = "(]"
//输出：false

func isValid(s string) bool {
	stack := &Stack{}
	for i := 0; i < len(s); i++ {
		v := s[i]
		switch v {
		case ')', ']', '}':
			data := stack.pop()
			if v == ')' && data != '(' {
				return false
			} else if v == ']' && data != '[' {
				return false
			} else if v == '}' && data != '{' {
				return false
			}
		default:
			stack.push(v)
		}
	}
	return stack.isEmpty()
}

type Stack struct {
	element []byte
}

func (s *Stack) push(data byte) {
	s.element = append(s.element, data)
}

func (s *Stack) pop() byte {
	if len(s.element) == 0 {
		return '-'
	}
	data := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]
	return data
}

func (s *Stack) isEmpty() bool {
	return len(s.element) == 0
}
