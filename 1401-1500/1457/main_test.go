package _1457

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, pseudoPalindromicPaths(&TreeNode{
		2,
		&TreeNode{
			3,
			&TreeNode{3, nil, nil},
			&TreeNode{1, nil, nil}},
		&TreeNode{1, nil, &TreeNode{1, nil, nil}}}))
}
