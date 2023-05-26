package _120

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(11, minimumTotal2([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
