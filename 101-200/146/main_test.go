package _46

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(2, 2)
	a.Equal(2, lru.Get(2))
	lru.Put(1, 1)
	lru.Put(4, 1)
	a.Equal(-1, lru.Get(2))
	//a.Equal(-1, lru.Get(1))
	//a.Equal(3, lru.Get(3))
	//a.Equal(4, lru.Get(4))
}
