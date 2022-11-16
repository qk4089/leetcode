package _25

//给你链表的头节点head，每k个节点一组进行翻转，请你返回修改后的链表。
//k是一个正整数，它的值小于或等于链表的长度。如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

//提示：
//	链表中的节点数目为 n
//	1 <= k <= n <= 5000
//	0 <= Node.val <= 1000

type ListNode struct {
	Val  int
	Next *ListNode
}

func main(head *ListNode, k int) *ListNode {
	p1, p2 := head, head
	count := 0
	for ; count < k && p1 != nil; count++ {
		p2, p1 = p1, p1.Next
	}
	if count < k {
		return head
	}
	p2.Next = nil
	result := reverse(head)
	head.Next = main(p1, k)
	return result
}

func reverse(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		tmp := head
		head, tmp.Next, result = head.Next, result, tmp
	}
	return result
}
