package _234

//给你一个单链表的头节点head，请你判断该链表是否为回文链表。如果是，返回true；否则，返回false。

//提示：
//	链表中节点数目在范围[1, 105] 内
//	0 <= Node.val <= 9

//进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var result *ListNode
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow, slow.Next, result = slow.Next, result, slow
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if result.Val != slow.Val {
			return false
		}
		result, slow = result.Next, slow.Next
	}
	return true
}
