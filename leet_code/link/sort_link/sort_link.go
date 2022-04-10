package sort_link

import (
	. "github.com/aozeahj/go_algorithm/leet_code/link"
)

/**
148. 排序链表

给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/

func merge(head1, head2 *ListNode) *ListNode {
	newHead := new(ListNode)
	cur, cur1, cur2 := newHead, head1, head2
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			cur.Next, cur1 = cur1, cur1.Next
		} else {
			cur.Next, cur2 = cur2, cur2.Next
		}
		cur = cur.Next
	}

	if cur1 != nil {
		cur.Next = cur1
	} else if cur2 != nil {
		cur.Next = cur2
	}

	return newHead.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	l := 0
	for cur != nil {
		cur = cur.Next
		l++
	}

	newHead := new(ListNode)
	newHead.Next = head
	for i := 1; i < l; i <<= 1 {
		prev := newHead
		cur = newHead.Next
		for cur != nil {
			head1 := cur
			for j := 1; j < i && cur.Next != nil; j++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for j := 1; j < i && cur != nil && cur.Next != nil; j++ {
				cur = cur.Next
			}

			if cur != nil {
				cur, cur.Next = cur.Next, nil
			}

			prev.Next = merge(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
		}
	}

	return newHead.Next
}
