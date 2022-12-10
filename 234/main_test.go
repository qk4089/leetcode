package _234

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	l1 := &ListNode{1, nil}
	a.True(isPalindrome(l1))
	l2 := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	a.True(isPalindrome(l2))
	l3 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	a.True(isPalindrome(l3))
}
