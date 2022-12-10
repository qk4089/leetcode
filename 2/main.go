package _2

//给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。你可以假设除了数字0之外，这两个数都不会以0开头。

//提示：
//	每个链表中的节点数在范围[1,100]内
//	0<=Node.val<=9
//	题目数据保证列表表示的数字不含前导零

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	v1, v2, overflow, sum := 0, 0, 0, 0
	root := &ListNode{-1, nil}
	point := root
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		sum = v1 + v2 + overflow
		overflow = sum / 10
		sum = sum % 10
		point.Next = &ListNode{sum % 10, nil}
		point = point.Next
	}
	if overflow > 0 {
		point.Next = &ListNode{1, nil}
	}
	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
