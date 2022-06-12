package _10_remove_nth_node_from_end_of_list

/**
19. 删除链表的倒数第 N 个结点

给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//计算链表长度
	l := 0
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}

	k := l - n + 1

	//删除正数第一个元素之前的元素 || 删除的是最后一个元素之后的元素  这种情况非法，不做删除操作
	if k < 1 || k > l {
		return head
	}

	//第一个元素需要单独出来
	if k == 1 {
		return head.Next
	}

	// 第 2~l 个元素
	cur = head
	for i := 1; i < k-1; i++ {
		cur = cur.Next
	}
	cur.Next, cur.Next.Next = cur.Next.Next, nil

	return head
}
