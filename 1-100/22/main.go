package _22

//https://leetcode.cn/problems/generate-parentheses/
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

//示例 1：
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]

//输入：n = 1
//输出：["()"]

//提示：
//	1 <= n <= 8

func generateParenthesis(n int) []string {
	existMap, cache := make(map[string]bool), make(map[int][]string)
	return generate(n, existMap, cache)
}

func generate(n int, existMap map[string]bool, cache map[int][]string) []string {
	if v, ok := cache[n]; ok {
		return v
	}
	if n == 1 {
		cache[1], existMap["()"] = []string{"()"}, true
		return cache[1]
	}
	if n == 2 {
		cache[2], existMap["()()"], existMap["(())"] = []string{"()()", "(())"}, true, true
		return cache[2]
	}
	result := make([]string, 0)
	for _, str := range generate(n-1, existMap, cache) {
		prefix, mid, suffix := "()"+str, "("+str+")", str+"()"
		if !existMap[prefix] {
			result, existMap[prefix] = append(result, prefix), true
		}
		if !existMap[mid] {
			result, existMap[mid] = append(result, mid), true
		}
		if !existMap[suffix] {
			result, existMap[suffix] = append(result, suffix), true
		}
	}
	for i := 1; i < n-1; i++ {
		for _, x := range generate(i, existMap, cache) {
			for _, y := range generate(n-1-i, existMap, cache) {
				str := "(" + x + ")" + y
				if !existMap[str] {
					result, existMap[str] = append(result, str), true
				}
			}
		}
	}
	cache[n] = result
	return result
}
