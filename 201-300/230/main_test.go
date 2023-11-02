package _230

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	a := assert.New(t)
	root := &TreeNode{3,
		&TreeNode{1, nil, &TreeNode{2, nil, nil}},
		&TreeNode{4, nil, nil},
	}
	a.Equal(1, kthSmallest(root, 1))
}
