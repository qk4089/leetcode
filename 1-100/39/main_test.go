package _39

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.ElementsMatch([][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8))
}
