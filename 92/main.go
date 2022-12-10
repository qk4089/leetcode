package _92

//给你单链表的头指针head和两个整数left和right，其中left<=right。请你反转从位置left到位置right的链表节点，返回反转后的链表。
//
//提示：
//	链表中节点数目为n
//	1<=n<=500
//	-500<=Node.val<=500
//	1<=left<=right<=n

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil || left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	prefix := dummy
	for i := 1; i < left; i++ {
		prefix = head
		head = head.Next
	}
	var result *ListNode
	suffix := head
	for i := 0; i <= right-left; i++ {
		tmp := head
		head, tmp.Next = head.Next, result
		result = tmp
	}
	prefix.Next = result
	suffix.Next = head
	return dummy.Next
}
