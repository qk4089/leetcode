package _206

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//提示：
//	链表中节点的数目范围是 [0, 5000]
//	-5000 <= Node.val <= 5000

type ListNode struct {
	Val  int
	Next *ListNode
}

func main(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		tmp := head
		head, tmp.Next = head.Next, result
		result = tmp
	}
	return result
}
