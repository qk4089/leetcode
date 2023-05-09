package _206

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_d(t *testing.T) {
	a := assert.New(t)
	a.Nil(reverse_d(nil))
	l1 := &ListNode{1, nil}
	a.Equal(l1, reverse_d(l1))
	l2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printNode(reverse_d(l2))
}

func TestN_d(t *testing.T) {
	a := assert.New(t)
	//a.Nil(reverseN_d(nil, 2))
	l1 := &ListNode{1, nil}
	a.Equal(l1, reverseN_d(l1, 2))
	l2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printNode(reverseN_d(l2, 3))
}
