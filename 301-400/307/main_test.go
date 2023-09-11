package _307

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	numArray := Constructor([]int{1, 3, 5})
	a.Equal(9, numArray.SumRange(0, 2))
	numArray.Update(1, 2)
	a.Equal(8, numArray.SumRange(0, 2))
}
