package _102

//https://leetcode.cn/problems/binary-tree-level-order-traversal/
//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

//示例 1：
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]

//输入：root = [1]
//输出：[[1]]

//输入：root = []
//输出：[]

//提示：
//	树中节点数目在范围 [0, 2000] 内
//	-1000 <= Node.val <= 1000

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmp, size := make([]int, len(queue)), len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			tmp[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans, queue = append(ans, tmp), queue[size:]
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
