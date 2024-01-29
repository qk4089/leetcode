package _416

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.True(canPartition([]int{1, 5, 11, 5}))
}
