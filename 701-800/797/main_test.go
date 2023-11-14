package _797

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{0, 1, 3}, {0, 2, 3}}, allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}))
}
