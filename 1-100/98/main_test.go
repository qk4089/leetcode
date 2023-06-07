package _98

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	a.False(isValidBST(&TreeNode{5, &TreeNode{4, nil, nil}, &TreeNode{6, &TreeNode{3., nil, nil}, &TreeNode{7, nil, nil}}}))
}
