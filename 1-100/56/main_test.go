package _56

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{1, 4}, {5, 6}}, merge([][]int{{1, 4}, {5, 6}}))
	a.Equal([][]int{{1, 6}, {8, 10}, {15, 18}}, merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
