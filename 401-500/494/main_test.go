package _494

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_51(t *testing.T) {
	a := assert.New(t)
	a.Equal(256, findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
