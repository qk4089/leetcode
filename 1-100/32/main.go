package _32

//https://leetcode.cn/problems/longest-valid-parentheses/description/
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

//示例 1：
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"

//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"

//输入：s = ""
//输出：0

//提示：
//	0 <= s.length <= 3 * 10^4
//	s[i] 为 '(' 或 ')'

func longestValidParentheses(s string) int {
	return dp(s)
}

func dp(s string) int {
	ans, result := 0, make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			if i >= 2 {
				result[i] = result[i-2] + 2
			} else {
				result[i] = 2
			}
		} else if i-result[i-1] >= 1 && s[i-result[i-1]-1] == '(' {
			if i-result[i-1] >= 2 {
				result[i] = result[i-1] + result[i-result[i-1]-2] + 2
			} else {
				result[i] = result[i-1] + 2
			}
		}
		ans = max(ans, result[i])
	}
	return ans
}

func stack(s string) int {
	ans, stk := 0, []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk = append(stk, i)
		} else {
			stk = stk[:len(stk)-1]
			if len(stk) == 0 {
				stk = append(stk, i)
			} else {
				ans = max(ans, i-stk[len(stk)-1])
			}
		}
	}
	return ans
}

func point(s string) int {
	ans, left, right := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			ans = max(ans, 2*right)
		} else if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			ans = max(ans, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
