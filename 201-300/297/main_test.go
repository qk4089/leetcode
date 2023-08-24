package _297

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	codec := Constructor()
	fmt.Println(codec.deserialize(codec.serialize(&TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil}},
	})))
}
