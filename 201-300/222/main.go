package _222

import "math"

//https://leetcode.cn/problems/count-complete-tree-nodes/
//给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
//完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
//并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

//示例 1：
//输入：root = [1,2,3,4,5,6]
//输出：6

//输入：root = []
//输出：0

//输入：root = [1]
//输出：1

//提示：
//	树中节点的数目范围是[0, 5 * 10^4]
//	0 <= Node.val <= 5 * 10^4
//	题目数据保证输入的树是 完全二叉树

//进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, leftH, right, rightH := root.Left, 1, root.Right, 1
	for left != nil {
		left, leftH = left.Left, leftH+1
	}
	for right != nil {
		right, rightH = right.Right, rightH+1
	}
	if leftH == rightH {
		return int(math.Pow(2.0, float64(leftH))) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
