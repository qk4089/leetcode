package _64

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(7, minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	a.Equal(12, minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
