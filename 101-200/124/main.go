package _124

//二叉树中的路径被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。
//该路径 至少包含一个 节点，且不一定经过根节点。
//
//路径和是路径中各节点值的总和。
//给你一个二叉树的根节点 root ，返回其 最大路径和 。

//示例 1：
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42

//提示：
//	树中节点数目范围是 [1, 3 * 10^4]
//	-1000 <= Node.val <= 1000

func maxPathSum(root *TreeNode) int {
	_, maxVal := loop(root)
	return maxVal
}

func loop(root *TreeNode) (int /*sideMax*/, int /*Max*/) {
	if root == nil {
		return -1001, -1001
	}
	leftSideMax, leftMax := loop(root.Left)
	rightSideMax, rightMax := loop(root.Right)
	sideMax := max(root.Val, root.Val+max(leftSideMax, rightSideMax))
	return sideMax, max(max(sideMax, root.Val+leftSideMax+rightSideMax), max(leftMax, rightMax))
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
