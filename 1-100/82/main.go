package _82

//https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
//给定一个已排序的链表的头head，删除原始链表中所有重复数字的节点，只留下不同的数字。返回已排序的链表

//示例
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]

//输入：head = [1,1,1,2,3]
//输出：[2,3]

//提示：
//	链表中节点数目在范围 [0, 300] 内
//	-100 <= Node.val <= 100
//	题目数据保证链表已经按升序 排列

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root := &ListNode{0, head}
	for point := root; point.Next != nil && point.Next.Next != nil; {
		if point.Next.Val != point.Next.Next.Val {
			point = point.Next
		} else {
			x := point.Next.Val
			for point.Next != nil && x == point.Next.Val {
				point.Next = point.Next.Next
			}
		}
	}
	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
