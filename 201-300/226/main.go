package _226

//给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

//示例
//输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]

//输入：root = [2,1,3]
//输出：[2,3,1]

//输入：root = []
//输出：[]

//提示：
//	树中节点数目范围在 [0, 100] 内
//	-100 <= Node.val <= 100

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
