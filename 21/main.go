package _21

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

type ListNode struct {
	Val  int
	Next *ListNode
}

func main(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	point := result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			point.Next = list1
			list1 = list1.Next
		} else {
			point.Next = list2
			list2 = list2.Next
		}
		point = point.Next
	}
	if list1 != nil {
		point.Next = list1
	} else {
		point.Next = list2
	}
	return result.Next
}
