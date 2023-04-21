package _206

// 使用递归实现
func reverse_d(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := reverse_d(head.Next)
	head.Next.Next = head
	head.Next = nil
	return result
}

// 反转前N个节点
var next *ListNode

func reverseN_d(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if n == 1 {
		next = head.Next
		return head
	}
	result := reverseN_d(head.Next, n-1)
	head.Next.Next = head
	head.Next = next
	return result
}
