package _112

//给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在根节点到叶子节点的路径，
//这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
//叶子节点 是指没有子节点的节点。

//示例
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
//输出：true
//解释：等于目标和的根节点到叶节点路径如上图所示。

//输入：root = [1,2,3], targetSum = 5
//输出：false
//解释：树中存在两条根节点到叶子节点的路径：
//(1 --> 2): 和为 3
//(1 --> 3): 和为 4
//不存在 sum = 5 的根节点到叶子节点的路径。

//输入：root = [], targetSum = 0
//输出：false
//解释：由于树是空的，所以不存在根节点到叶子节点的路径。

//提示：
//	树中节点的数目在范围 [0, 5000] 内
//	-1000 <= Node.val <= 1000
//	-1000 <= targetSum <= 1000

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queue := []*Node{{root, 0}}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := pop(&queue)
			sum := node.sum + node.treeNode.Val
			if node.treeNode.Left == nil && node.treeNode.Right == nil && sum == targetSum {
				return true
			}
			if node.treeNode.Left != nil {
				queue = append(queue, &Node{node.treeNode.Left, sum})
			}
			if node.treeNode.Right != nil {
				queue = append(queue, &Node{node.treeNode.Right, sum})
			}
		}
		size = len(queue)
	}
	return false
}

func pop(queue *[]*Node) *Node {
	node := (*queue)[0]
	*queue = (*queue)[1:]
	return node
}

type Node struct {
	treeNode *TreeNode
	sum      int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
