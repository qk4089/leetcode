package _83

//给定一个已排序的链表的头head，删除所有重复的元素，使每个元素只出现一次。返回已排序的链表。

//提示：
//	链表中节点数目在范围[0,300]内
//	-100<=Node.val<=100
//	题目数据保证链表已经按升序排列

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		index, point := head, head.Next
		for point != nil {
			if point.Val != index.Val {
				index.Next, index = point, point
			}
			point = point.Next
		}
		index.Next = nil
	}
	return head
}
