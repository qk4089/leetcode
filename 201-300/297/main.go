package _297

import (
	"strconv"
	"strings"
)

//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输
//到另一个计算机环境，采取相反方式重构得到原数据。
//请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列/反序列化算法执行逻辑， 你只需要保证一个二叉树可以被序列化 为一个字符串
//并且将这个字符串反序列化为原始的树结构。

//提示: 输入输出格式与LeetCode目前使用的方式一致，详情请参阅LeetCode序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

//示例 1：

//输入：root = [1,2,3,null,null,4,5]
//输出：[1,2,3,null,null,4,5]

//输入：root = []
//输出：[]

//输入：root = [1]
//输出：[1]

//输入：root = [1,2]
//输出：[1,2]

//提示：
//	树中结点数在范围 [0, 10^4] 内
//	-1000 <= Node.val <= 1000

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	return build(&nodes)
}

func build(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	first := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if first == "#" {
		return nil
	}
	val, _ := strconv.Atoi(first)
	parent := &TreeNode{val, build(nodes), build(nodes)}
	return parent
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
