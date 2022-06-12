package _12_merge_two_sorted_lists

/**

21. 合并两个有序链表

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := new(ListNode)
	cur := newHead
	cur1 := list1
	cur2 := list2

	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next, cur1 = cur1, cur1.Next
		} else {
			cur.Next, cur2 = cur2, cur2.Next
		}

		cur = cur.Next
	}

	if cur1 != nil {
		cur.Next = cur1
	}

	if cur2 != nil {
		cur.Next = cur2
	}

	return newHead.Next
}
