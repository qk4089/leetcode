package _347

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.ElementsMatch([]int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
