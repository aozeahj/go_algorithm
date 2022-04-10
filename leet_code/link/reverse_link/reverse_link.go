package reverse_link

import . "github.com/aozeahj/go_algorithm/leet_code/link"

/**
206. 反转链表

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
*/
func reverseList(head *ListNode) *ListNode {
	newHead := new(ListNode)
	curHead := head
	for curHead != nil {
		newHead.Next, curHead.Next, curHead = curHead, newHead.Next, curHead.Next
	}

	return newHead.Next
}
