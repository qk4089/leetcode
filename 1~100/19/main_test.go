package _19

import "testing"

func Test(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, nil}}
	removeNthFromEnd(l1, 2)
	//l2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//removeNthFromEnd(l2, 2)
}
