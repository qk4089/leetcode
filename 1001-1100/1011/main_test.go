package _1011

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(15, shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	a.Equal(6, shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	a.Equal(3, shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
}
