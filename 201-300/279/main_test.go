package _279

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, numSquares(12))
	a.Equal(2, numSquares(13))
}
