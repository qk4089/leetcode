package _19

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

//提示：
//	链表中结点的数目为 sz
//	1 <= sz <= 30
//	0 <= Node.val <= 100
//	1 <= n <= sz

type ListNode struct {
	Val  int
	Next *ListNode
}

func main(head *ListNode, n int) *ListNode {
	slow, point := head, head
	for i := 0; i <= n; i++ {
		if point == nil {
			return slow.Next
		}
		point = point.Next
	}
	for point != nil {
		point = point.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
