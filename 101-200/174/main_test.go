package _174

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, calculateMinimumHP([][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}))
	a.Equal(7, calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
	a.Equal(3, calculateMinimumHP([][]int{{-2}}))
	a.Equal(1, calculateMinimumHP([][]int{{100}}))
}
