package _07

//https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
//给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。（即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

//示例 1：
//输入：root = [3,9,20,null,null,15,7]
//输出：[[15,7],[9,20],[3]]

//输入：root = [1]
//输出：[[1]]

//输入：root = []
//输出：[]

//提示：
//	树中节点数目在范围 [0, 2000] 内
//	-1000 <= Node.val <= 1000

func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmp, size := make([]int, 0), len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans, queue = append(ans, tmp), queue[size:]
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
