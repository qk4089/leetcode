package _652

import (
	"testing"
)

func Test(t *testing.T) {
	findDuplicateSubtrees(&TreeNode{10,
		&TreeNode{2,
			&TreeNode{1, nil, nil},
			&TreeNode{12, nil, nil}},
		&TreeNode{22,
			&TreeNode{1, nil, nil},
			&TreeNode{1, nil, nil}},
	})
}
