package _337

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.Equal(7, rob(&TreeNode{3,
		&TreeNode{2, nil, &TreeNode{3, nil, nil}},
		&TreeNode{3, nil, &TreeNode{1, nil, nil}}},
	))
	a.Equal(7, rob(&TreeNode{4,
		&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil},
		nil},
	))
}
