package _78

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.ElementsMatch([][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}, subsets([]int{1, 2, 3}))
	a.ElementsMatch([][]int{{}, {3, 2, 4, 1}, {3}, {2}, {4}, {1}, {3, 2}, {3, 4}, {3, 1}, {3, 2, 4}, {3, 2, 1}, {3, 4, 1}, {2, 4}, {2, 1}, {2, 4, 1}, {4, 1}}, subsets([]int{3, 2, 4, 1}))
}
