package _48

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	fmt.Println(sortList(head))
}
