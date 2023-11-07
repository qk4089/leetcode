package _95

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(5, len(generateTrees(3)))
	a.Equal(1, len(generateTrees(1)))
	a.Equal(14, len(generateTrees(4)))
}
