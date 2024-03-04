package _138

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	node1 := &Node{3, nil, nil}
	node2 := &Node{3, nil, nil}
	node3 := &Node{3, nil, nil}
	node4 := &Node{10, nil, nil}
	node5 := &Node{1, nil, nil}
	node1.Next = node2
	node2.Next = node3
	node2.Random = node1
	node3.Next = node4
	node3.Random = node5
	node4.Next = node5
	node4.Random = node3
	node5.Random = node1
	fmt.Println(copyRandomList(node1))
}
