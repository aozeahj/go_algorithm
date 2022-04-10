package swap_nodes_in_pairs

import . "github.com/aozeahj/go_algorithm/leet_code/link"

/**
24. 两两交换链表中的节点

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

输入：head = [1,2,3,4]
输出：[2,1,4,3]
*/
func swapPairs(head *ListNode) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	prevNode := newHead
	for prevNode.Next != nil && prevNode.Next.Next != nil {
		prevNode.Next, prevNode.Next.Next.Next, prevNode.Next.Next = prevNode.Next.Next, prevNode.Next, prevNode.Next.Next.Next
		prevNode = prevNode.Next.Next
	}

	return newHead.Next
}
