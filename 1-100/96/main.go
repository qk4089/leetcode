package _96

import "strconv"

//https://leetcode.cn/problems/unique-binary-search-trees/
//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

//示例 1：
//
//输入：n = 3
//输出：5

//输入：n = 1
//输出：1
//提示：
//	1 <= n <= 19

func numTrees(n int) int {
	areaMap := make(map[string]int)
	return count(1, n, areaMap)
}

func count(l, r int, areaMap map[string]int) int {
	if l > r {
		return 1
	}
	tmp := strconv.Itoa(l) + "+" + strconv.Itoa(r)
	if val, exist := areaMap[tmp]; exist {
		return val
	}
	ans := 0
	for i := l; i <= r; i++ {
		left := count(l, i-1, areaMap)
		right := count(i+1, r, areaMap)
		ans += left * right
	}
	areaMap[tmp] = ans
	return ans
}
