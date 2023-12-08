package _144

//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

//示例 1：
//输入：root = [1,null,2,3]
//输出：[1,2,3]

//输入：root = []
//输出：[]

//输入：root = [1]
//输出：[1]

//输入：root = [1,2]
//输出：[1,2]

//输入：root = [1,null,2]
//输出：[1,2]

//提示：
//	树中节点数目在范围 [0, 100] 内
//	-100 <= Node.val <= 100
//进阶：递归算法很简单，你可以通过迭代算法完成吗？

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue, result = queue[:len(queue)-1], append(result, node.Val)
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
