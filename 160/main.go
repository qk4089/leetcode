package _160

//给你两个单链表的头节点headA和headB，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回null。
//题目数据 保证 整个链式结构中不存在环。注意，函数返回结果后，链表必须 保持其原始结构 。

//提示：
//	listA 中节点数目为 m
//	listB 中节点数目为 n
//	1 <= m, n <= 3 * 104
//	1 <= Node.val <= 105

//进阶：你能否设计一个时间复杂度 O(m + n) 、仅用 O(1) 内存的解决方案？

type ListNode struct {
	Val  int
	Next *ListNode
}

func main(headA, headB *ListNode) *ListNode {
	p0, p1 := headA, headB
	for p0 != p1 {
		if p0 == nil {
			p0 = headB
		} else {
			p0 = p0.Next
		}
		if p1 == nil {
			p1 = headA
		} else {
			p1 = p1.Next
		}
	}
	return p0
}
