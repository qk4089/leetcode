package _206

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//提示：
//	链表中节点的数目范围是 [0, 5000]
//	-5000 <= Node.val <= 5000

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		head, head.Next, result = head.Next, result, head
	}
	return result
}

// 反转前N个节点(包含N)
func reverseN(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	var result *ListNode
	point := head
	for count := 0; point != nil && count < n; count++ {
		tmp := point
		point, tmp.Next, result = point.Next, result, tmp
	}
	head.Next = point
	return result
}
