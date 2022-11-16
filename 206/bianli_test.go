package _206

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Nil(reverse(nil))
	l1 := &ListNode{1, nil}
	a.Equal(l1, reverse(l1))
	l2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printNode(reverse(l2))
}

func TestN(t *testing.T) {
	a := assert.New(t)
	a.Nil(reverseN(nil, 2))
	l1 := &ListNode{1, nil}
	a.Equal(l1, reverseN(l1, 2))
	l2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printNode(reverseN(l2, 3))
}

func printNode(head *ListNode) {
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
}
