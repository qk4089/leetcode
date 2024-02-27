package _199

//https://leetcode.cn/problems/binary-tree-right-side-view/
//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

//示例 1:
//输入: [1,2,3,null,5,null,4]
//输出: [1,3,4]

//输入: [1,null,3]
//输出: [1,3]

//输入: []
//输出: []

//提示:
//二叉树的节点个数的范围是 [0,100]
//-100 <= Node.val <= 100

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans, queue := make([]int, 0), []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i == size-1 {
				ans = append(ans, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
