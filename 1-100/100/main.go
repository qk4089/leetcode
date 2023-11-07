package _100

//https://leetcode.cn/problems/same-tree/description/
//给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

//示例 1：
//输入：p = [1,2,3], q = [1,2,3]
//输出：true

//输入：p = [1,2], q = [1,null,2]
//输出：false

//输入：p = [1,2,1], q = [1,1,2]
//输出：false

//提示：
//	两棵树上的节点数目都在范围 [0, 100] 内
//	-10^4 <= Node.val <= 10^4

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
