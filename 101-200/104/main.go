package _104

//给定一个二叉树，找出其最大深度。二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//说明: 叶子节点是指没有子节点的节点。
//
//示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//		3
//	   / \
//	  9  20
//      /  \
//	   15   7
//返回它的最大深度 3 。

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := maxDepth(root.Left), maxDepth(root.Right)
	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
