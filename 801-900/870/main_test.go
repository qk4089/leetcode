package _870

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	//a.Equal([]int{2, 11, 7, 15}, advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	//a.Equal([]int{2, 0, 2, 1, 4}, advantageCount([]int{2, 0, 4, 1, 2}, []int{1, 3, 0, 0, 2}))
	a.Equal([]int{15448, 13574, 6475, 19893, 14234}, advantageCount([]int{15448, 14234, 13574, 19893, 6475}, []int{14234, 6475, 19893, 15448, 13574}))
}
