package _155

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	minStack := Constructor()
	minStack.Push(2)
	minStack.Push(1)
	minStack.Push(3)
	minStack.Push(4)
	a.Equal(4, minStack.Top())
	a.Equal(1, minStack.GetMin())
	minStack.Pop()
	a.Equal(3, minStack.Top())
	a.Equal(1, minStack.GetMin())
	minStack.Pop()
	a.Equal(1, minStack.Top())
	a.Equal(1, minStack.GetMin())
	minStack.Pop()
	a.Equal(2, minStack.Top())
	a.Equal(2, minStack.GetMin())
	minStack.Push(4)
	a.Equal(4, minStack.Top())
	a.Equal(2, minStack.GetMin())
	minStack.Push(1)
	a.Equal(1, minStack.Top())
	a.Equal(1, minStack.GetMin())
}
