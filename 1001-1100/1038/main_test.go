package _1038

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	root := &TreeNode{3,
		&TreeNode{1, nil, &TreeNode{2, nil, nil}},
		&TreeNode{4, nil, nil},
	}
	fmt.Println(bstToGst(root))
}
