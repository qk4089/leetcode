package _160

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//l := &ListNode{2, &ListNode{4, nil}}
	l1 := &ListNode{2, &ListNode{6, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{5, nil}}
	fmt.Println(getIntersectionNode(l1, l2))
}
