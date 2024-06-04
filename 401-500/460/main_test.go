package _460

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	a.Equal(1, lfu.Get(1))
	lfu.Put(3, 3)
	a.Equal(-1, lfu.Get(2))
	a.Equal(3, lfu.Get(3))
	lfu.Put(4, 4)
	a.Equal(-1, lfu.Get(1))
	a.Equal(3, lfu.Get(3))
	a.Equal(4, lfu.Get(4))
}
