package _89

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.ElementsMatch([]int{0, 1}, grayCode(1))
	//a.ElementsMatch([]int{0, 3, 2, 1}, grayCode(2))
	a.ElementsMatch([]int{0, 1, 3, 2, 6, 7, 5, 4}, grayCode(3))
}
