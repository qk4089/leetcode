package _236

import (
	"testing"
)

func Test(t *testing.T) {
	a := &TreeNode{7, nil, nil}
	b := &TreeNode{11, nil, nil}
	root := &TreeNode{3, &TreeNode{4, a, a}, &TreeNode{5, nil, nil}}
	lowestCommonAncestor(root, a, b)
}
