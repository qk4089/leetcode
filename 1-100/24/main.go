package _24

//https://leetcode.cn/problems/swap-nodes-in-pairs/
//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

//示例 1：
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]

//输入：head = []
//输出：[]

//输入：head = [1]
//输出：[1]

//提示：
//	链表中节点的数目在范围 [0, 100] 内
//	0 <= Node.val <= 100

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node, next := swapPairs(head.Next.Next), head.Next
	head.Next, next.Next = node, head
	return next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
