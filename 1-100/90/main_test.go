package _90

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.ElementsMatch([][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}, subsetsWithDup([]int{1, 2, 3}))
	a.ElementsMatch([][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}, subsetsWithDup([]int{1, 2, 2}))
	a.ElementsMatch([][]int{{}, {1}, {2}, {1, 1}, {1, 2}, {1, 1, 2}}, subsetsWithDup([]int{1, 1, 2}))
}
