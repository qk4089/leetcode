package _450

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	root := &TreeNode{5,
		&TreeNode{3,
			&TreeNode{2, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{6,
			nil,
			&TreeNode{7, nil, nil}},
	}
	node := deleteNode(root, 3)
	fmt.Println(node)
}
