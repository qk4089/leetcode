package _652

import "strconv"

//给你一棵二叉树的根节点 root ，返回所有重复的子树。 对于同一类的重复子树，你只需要返回其中任意 一棵 的根结点即可。
//如果两棵树具有相同的结构和相同的结点值，则认为二者是重复的。

//示例 1：
//输入：root = [1,2,3,4,null,2,4,null,null,4]
//输出：[[2,4],[4]]

//输入：root = [2,1,1]
//输出：[[1]]

//输入：root = [2,2,2,3,null,3,null]
//输出：[[2,3],[3]]

//提示：
//	树中的结点数在 [1, 5000] 范围内。
//	-200 <= Node.val <= 200

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	result, nodeMap := make([]*TreeNode, 0), make(map[string]int)
	dfs(root, nodeMap, &result)
	return result
}

func dfs(node *TreeNode, nodeMap map[string]int, result *[]*TreeNode) string {
	if node == nil {
		return "#"
	}
	left, right := dfs(node.Left, nodeMap, result), dfs(node.Right, nodeMap, result)
	tmp := left + right + strconv.Itoa(node.Val) + ","
	if val, exist := nodeMap[tmp]; exist {
		if val == 1 {
			*result = append(*result, node)
		}
	}
	nodeMap[tmp]++
	return tmp
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
