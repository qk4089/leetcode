package _230

//https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

//示例 1：
//输入：root = [3,1,4,null,2], k = 1
//输出：1

//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3

//提示：
//	树中的节点数为 n 。
//	1 <= k <= n <= 10^4
//	0 <= Node.val <= 10^4

//进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？

func kthSmallest(root *TreeNode, k int) int {
	cnt, res := 0, 0
	traverse(root, k, &cnt, &res)
	return res
}

func traverse(root *TreeNode, k int, cnt, res *int) {
	if root == nil {
		return
	}
	traverse(root.Left, k, cnt, res)
	*cnt++
	if *cnt == k {
		*res = root.Val
		return
	}
	traverse(root.Right, k, cnt, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
