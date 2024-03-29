package _106

//给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。

//示例 1:
//输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
//输出：[3,9,20,null,null,15,7]

//输入：inorder = [-1], postorder = [-1]
//输出：[-1]

//提示:
//	1 <= inorder.length <= 3000
//	postorder.length == inorder.length
//	-3000 <= inorder[i], postorder[i] <= 3000
//	inorder 和 postorder 都由 不同 的值组成
//	postorder 中每一个值都在 inorder 中
//	inorder 保证是树的中序遍历
//	postorder 保证是树的后序遍历

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	idx, val := 0, postorder[len(postorder)-1]
	node := &TreeNode{Val: val}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			idx = i
			break
		}
	}
	if idx > 0 {
		node.Left = buildTree(inorder[:idx], postorder[:idx])
	}
	if idx < len(inorder)-1 {
		node.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	}
	return node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
