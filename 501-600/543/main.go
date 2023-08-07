package _543

//给你一棵二叉树的根节点，返回该树的直径 。
//二叉树的直径是指树中任意两个节点之间最长路径的长度 。这条路径可能经过也可能不经过根节点root。
//两节点之间路径的长度由它们之间边数表示。

//示例 1：
//输入：root = [1,2,3,4,5]
//输出：3
//解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。

//输入：root = [1,2]
//输出：1

//提示：
//	树中节点数目在范围 [1, 10^4] 内
//	-100 <= Node.val <= 100

func diameterOfBinaryTree(root *TreeNode) int {
	_, length := getMax(root)
	return length
}

func getMax(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftDepth, leftLength := getMax(root.Left)
	rightDepth, rightLength := getMax(root.Right)
	return max(leftDepth, rightDepth) + 1, max(leftDepth+rightDepth, max(leftLength, rightLength))
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
