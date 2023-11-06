package _450

//https://leetcode.cn/problems/delete-node-in-a-bst/description/
//给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
//一般来说，删除节点可分为两个步骤：首先找到需要删除的节点；如果找到了，删除它。

//示例 1:

//输入：root = [5,3,6,2,4,null,7], key = 3
//输出：[5,4,6,2,null,null,7]
//解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
//一个正确的答案是 [5,4,6,2,null,null,7],
//另一个正确答案是 [5,2,6,null,4,null,7]。

//输入: root = [5,3,6,2,4,null,7], key = 0
//输出: [5,3,6,2,4,null,7]
//解释: 二叉树不包含值为 0 的节点

//输入: root = [], key = 0
//输出: []

//提示:
//	节点数的范围 [0, 10^4].
//	-10^5 <= Node.val <= 10^5
//	节点值唯一
//	root 是合法的二叉搜索树
//	-10^5 <= key <= 10^5

//进阶： 要求算法时间复杂度为 O(h)，h 为树的高度。

func deleteNode(root *TreeNode, key int) *TreeNode {
	sentry := &TreeNode{100001, root, nil}
	doDelete(sentry, key)
	return sentry.Left
}

func doDelete(node *TreeNode, key int) *TreeNode {
	var target *TreeNode
	isLeft := true
	if node == nil || node.Val == key {
		return node
	} else if node.Val > key {
		target = doDelete(node.Left, key)
	} else {
		target, isLeft = doDelete(node.Right, key), false
	}
	if target != nil {
		var replace *TreeNode
		if target.Left == nil {
			replace = target.Right
		} else if target.Right == nil {
			replace = target.Left
		} else {
			if target.Left.Right == nil {
				replace = target.Left
				target.Left.Right = target.Right
			} else {
				replace = findLeftMaxNode(target.Left.Right, target.Left)
				replace.Left, replace.Right = target.Left, target.Right
			}
		}
		if isLeft {
			node.Left = replace
		} else {
			node.Right = replace
		}
	}
	return nil
}

func findLeftMaxNode(node, parent *TreeNode) *TreeNode {
	if node.Right != nil {
		return findLeftMaxNode(node.Right, node)
	} else {
		if node.Left != nil {
			parent.Right = node.Left
		} else {
			parent.Right = nil
		}
		return node
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
