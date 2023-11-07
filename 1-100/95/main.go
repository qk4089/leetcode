package _95

//https://leetcode.cn/problems/unique-binary-search-trees-ii/
//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

//示例 1：

//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

//输入：n = 1
//输出：[[1]]

//提示：
//	1 <= n <= 8

func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func generate(l, r int) []*TreeNode {
	nodes := make([]*TreeNode, 0)
	if l > r {
		nodes = append(nodes, nil)
		return nodes
	}
	for i := l; i <= r; i++ {
		left := generate(l, i-1)
		right := generate(i+1, r)
		for x := 0; x < len(left); x++ {
			for y := 0; y < len(right); y++ {
				nodes = append(nodes, &TreeNode{i, left[x], right[y]})
			}
		}
	}
	return nodes
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
