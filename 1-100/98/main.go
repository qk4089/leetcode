package _98

import "math"

//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

//有效二叉搜索树定义如下：
//	节点的左子树只包含 小于 当前节点的数。
//	节点的右子树只包含 大于 当前节点的数。
//	所有左子树和右子树自身必须也是二叉搜索树。

//示例
//输入：root = [2,1,3]
//输出：true

//输入：root = [5,1,4,null,null,3,6]
//输出：false
//解释：根节点的值是 5 ，但是右子节点的值是 4 。

//提示：
//	树中节点数目范围在[1, 104] 内
//	-2^31 <= Node.val <= 2^31 - 1

func isValidBST(root *TreeNode) bool {
	return valid(root.Left, math.MinInt32, root.Val) && valid(root.Right, root.Val, math.MaxInt32)
}

func valid(node *TreeNode, low, high int) bool {
	if node == nil {
		return true
	}
	if node.Val > low && node.Val < high {
		return valid(node.Left, low, node.Val) && valid(node.Right, node.Val, high)
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
