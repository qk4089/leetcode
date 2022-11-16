package _25

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	l1 := &ListNode{1, nil}
	printNode(main(l1, 2))
	fmt.Println("============")
	l2 := &ListNode{1, &ListNode{2, nil}}
	printNode(main(l2, 2))
	fmt.Println("============")
	l3 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printNode(main(l3, 2))
}

func printNode(head *ListNode) {
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
}
