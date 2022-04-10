package reverse_k_nodes

import . "github.com/aozeahj/go_algorithm/leet_code/link"

/**
25. K 个一组翻转链表

给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	prev := newHead
	for head != nil {
		tail := head
		for i := 1; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return newHead.Next
			}
		}
		next := tail.Next
		newHead, newTail := reverse(head, tail)
		prev.Next = newHead
		newTail.Next = next
		prev = newTail
		head = prev.Next
	}
	return newHead.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	cur := head
	for cur != tail {
		tail.Next, cur.Next, cur = cur, tail.Next, cur.Next
	}

	return tail, head
}
