package _86

import (
	"testing"
)

func Test(t *testing.T) {
	//l1 := &ListNode{1, &ListNode{9, &ListNode{2, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, nil}}}}}}}
	//partition(l1, 3)
	l2 := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	partition(l2, 3)
}
