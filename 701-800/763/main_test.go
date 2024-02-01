package _763

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.ElementsMatch([]int{9, 7, 8}, partitionLabels("ababcbacadefegdehijhklij"))
}
