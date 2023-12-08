package _94

//https://leetcode.cn/problems/binary-tree-inorder-traversal/
//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

//示例 1：
//输入：root = [1,null,2,3]
//输出：[1,3,2]

//输入：root = []
//输出：[]

//输入：root = [1]
//输出：[1]

//提示：
//	树中节点数目在范围 [0, 100] 内
//	-100 <= Node.val <= 100

//进阶: 递归算法很简单，你可以通过迭代算法完成吗？

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		if node.Left != nil {
			queue, node.Left = append(queue, node.Left), nil
		} else {
			queue, ans = queue[:len(queue)-1], append(ans, node.Val)
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
