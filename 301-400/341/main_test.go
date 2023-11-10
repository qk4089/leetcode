package _341

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	result := make([]int, 0)
	iterator := Constructor([]*NestedInteger{
		//{0, 0, []*NestedInteger{}},
		{0, 0, []*NestedInteger{
			{1, 1, nil},
		}},
		{2, 1, nil},
		{4, 1, nil},
	})
	for iterator.HasNext() {
		result = append(result, iterator.Next())
	}
	a.Equal([]int{1, 3, 2, 4}, result)
}
