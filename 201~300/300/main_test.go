package _300

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal(4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//a.Equal(5, lengthOfLIS([]int{11, 1, 2, 9, 8, 4, 18, 22}))
	//a.Equal(3, lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	a.Equal(4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
}
