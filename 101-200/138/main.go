package _138

//https://leetcode.cn/problems/copy-list-with-random-pointer/
//给你一个长度为n的链表，每个节点包含一个额外增加的随机指针random，该指针可以指向链表中的任何节点或空节点。构造这个链表的深拷贝。深拷贝应该正好由
//n个全新节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的next指针和random指针也都应指向复制链表中的新节点，并使原链表和复制链表中的
//这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
//例如，如果原链表中有X和Y两个节点，其中X.random-->Y 。那么在复制链表中对应的两个节点x和y，同样有x.random-->y。返回复制链表的头节点。

//用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
//val：一个表示 Node.val 的整数。
//random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
//你的代码 只 接受原链表的头节点 head 作为传入参数。

//示例 1：
//输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
//输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

//输入：head = [[1,1],[2,1]]
//输出：[[1,1],[2,1]]

//输入：head = [[3,null],[3,0],[3,null]]
//输出：[[3,null],[3,0],[3,null]]

//提示：
//	0 <= n <= 1000
//	-10^4 <= Node.val <= 10^4
//	Node.random 为 null 或指向链表中的节点

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	var doGenerate func(node *Node) *Node
	doGenerate = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if val, exist := nodeMap[node]; exist {
			return val
		} else {
			nNode := &Node{node.Val, nil, nil}
			nodeMap[node] = nNode
			nNode.Next, nNode.Random = doGenerate(node.Next), doGenerate(node.Random)
			return nNode
		}
	}
	return doGenerate(head)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
