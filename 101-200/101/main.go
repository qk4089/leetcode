package _101

//https://leetcode.cn/problems/symmetric-tree/
//给你一个二叉树的根节点 root ， 检查它是否轴对称。

//示例 1：
//输入：root = [1,2,2,3,4,4,3]
//输出：true

//输入：root = [1,2,2,null,3,null,3]
//输出：false

//提示：
//	树中节点数目在范围 [1, 1000] 内
//	-100 <= Node.val <= 100

func isSymmetric(root *TreeNode) bool {
	return isSame(root.Left, root.Right)
}

func isSame(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 != nil && node2 == nil || node1 == nil && node2 != nil {
		return false
	}
	return node1.Val == node2.Val && isSame(node1.Left, node2.Right) && isSame(node1.Right, node2.Left)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
