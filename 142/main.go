package _142

//给定一个链表的头节点head，返回链表开始入环的第一个节点。如果链表无环，则返回null。
//如果链表中有某个节点，可以通过连续跟踪next指针再次到达，则链表中存在环。为了表示给定链表中的环，评测系统内部使用整数pos来表示链表尾连接到链表中的位置（索引从0开始）。
//如果pos是-1，则在该链表中没有环。注意：pos不作为参数进行传递，仅仅是为了标识链表的实际情况。不允许修改链表。

//提示：
//	链表中节点的数目范围在范围 [0, 104] 内
//	-105 <= Node.val <= 105
//	pos 的值为 -1 或者链表中的一个有效索引

//进阶：你是否可以使用 O(1) 空间解决此题？

type ListNode struct {
	Val  int
	Next *ListNode
}

func main(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}
	}
	return nil
}
