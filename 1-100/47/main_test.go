package _47

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.ElementsMatch([][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}, permuteUnique([]int{1, 1, 2}))
}
