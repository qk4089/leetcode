package _103

//https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

//示例 1：
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]

//输入：root = [1]
//输出：[[1]]

//输入：root = []
//输出：[]

//提示：
//	树中节点数目在范围 [0, 2000] 内
//	-100 <= Node.val <= 100

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	level, queue := 1, []*TreeNode{root}
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
		if level%2 == 0 {
			for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
		}
		level, ans, queue = level+1, append(ans, tmp), queue[size:]
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
