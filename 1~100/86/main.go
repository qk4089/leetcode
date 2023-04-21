package _86

//给你一个链表的头节点head和一个特定值x，请你对链表进行分隔，使得所有小于x的节点都出现在大于或等于x的节点之前。
//你应当保留两个分区中每个节点的初始相对位置

//提示：
//	链表中节点的数目在范围 [0, 200] 内
//	-100 <= Node.val <= 100
//	-200 <= x <= 200

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	l1, l2 := &ListNode{}, &ListNode{}
	p1, p2 := l1, l2
	for head != nil {
		tmp := head
		head, tmp.Next = head.Next, nil
		if tmp.Val < x {
			p1.Next = tmp
			p1 = p1.Next
		} else {
			p2.Next = tmp
			p2 = p2.Next
		}
	}
	p1.Next = l2.Next
	return l1.Next
}
