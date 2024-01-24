package _48

//https://leetcode.cn/problems/sort-list/
//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

//示例 1：
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]

//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]

//输入：head = []
//输出：[]

//提示：
//	链表中节点的数目在范围 [0, 5 * 10^4] 内
//	-10^5 <= Node.val <= 10^5

//进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

func sortList(head *ListNode) *ListNode {
	return merge(head)
}

func merge(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	slow, fast := node, node.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	next := slow.Next
	slow.Next = nil
	return contact(merge(node), merge(next))
}

func contact(node1, node2 *ListNode) *ListNode {
	root := &ListNode{-1, nil}
	point := root
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			point.Next, node1 = node1, node1.Next
		} else {
			point.Next, node2 = node2, node2.Next
		}
		point = point.Next
	}
	if node1 == nil {
		point.Next = node2
	} else {
		point.Next = node1
	}
	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
