package _23

//给你一个链表数组，每个链表都已经按升序排列。
//
//请你将所有链表合并到一个升序链表中，返回合并后的链表。

//提示：
//	k == lists.length
//	0 <= k <= 10^4
//	0 <= lists[i].length <= 500
//	-10^4 <= lists[i][j] <= 10^4
//	lists[i] 按 升序 排列
//	lists[i].length 的总和不超过 10^4

type ListNode struct {
	Val  int
	Next *ListNode
}

func main(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	return concat(merge(lists[:mid]), merge(lists[mid:]))
}

func concat(list1 *ListNode, list2 *ListNode) *ListNode {
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
