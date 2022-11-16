package _92

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	l1 := &ListNode{1, nil}
	printNode(main(l1, 1, 1))
	l2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printNode(main(l2, 1, 3))
	l3 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printNode(main(l3, 2, 4))
	l4 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printNode(main(l4, 1, 5))
}

func printNode(head *ListNode) {
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
}
