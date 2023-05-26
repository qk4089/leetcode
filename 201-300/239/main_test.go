package _239

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{1, -1}, maxSlidingWindow([]int{1, -1}, 1))
}
