package _96

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(5, numTrees(3))
	a.Equal(1, numTrees(1))
	a.Equal(14, numTrees(4))
}
