package _61

//https://leetcode.cn/problems/rotate-list/description/
//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

//示例 1：
//输入：head = [1,2,3,4,5], k = 2
//输出：[4,5,1,2,3]

//输入：head = [0,1,2], k = 4
//输出：[2,0,1]

//提示：
//链表中节点的数目在范围 [0, 500] 内
//	-100 <= Node.val <= 100
//	0 <= k <= 2 * 10^9

func rotateRight(head *ListNode, k int) *ListNode {
	point, sum, nodeMap := head, 0, make(map[int]*ListNode)
	for point != nil {
		nodeMap[sum], sum, point = point, sum+1, point.Next
	}
	if sum <= 1 {
		return head
	}
	cnt := k % sum
	for cnt > 0 {
		lastPreNode, lastNode := nodeMap[sum-2], nodeMap[sum-1]
		lastNode.Next, head, lastPreNode.Next = head, lastNode, nil
		cnt--
		sum--
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
