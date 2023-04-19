package _63

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(2, uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	//a.Equal(1, uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
	a.Equal(1, uniquePathsWithObstacles([][]int{{1}, {0}}))
}
