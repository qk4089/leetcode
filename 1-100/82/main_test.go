package _82

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(deleteDuplicates(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}))
}
