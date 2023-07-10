package _380

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	r := Constructor()
	a.True(r.Insert(2))
	a.True(r.Insert(0))
	r.GetRandom()
	a.True(r.Insert(-2))
	a.False(r.Insert(-2))
	a.False(r.Remove(3))
	a.False(r.Remove(3))
	a.False(r.Remove(1))
	a.True(r.Insert(-3))
	a.False(r.Remove(3))
	a.Equal(0, r.GetRandom())
}
