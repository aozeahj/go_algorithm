package _14_merge_k_sorted_lists

import (
	"container/heap"
	"image"
)

/**
23. 合并K个升序链表

给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type IntHeap []*ListNode

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	*h = old[:len(old)-1]
	return old[len(old)-1]
}

func mergeKLists(lists []*ListNode) *ListNode {
	priorityQueue := make(IntHeap, 0, len(lists))

	//将n个有序链表的头添加到优先队列中
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		priorityQueue = append(priorityQueue, lists[i])
	}

	//对优先级队列进行堆初始化
	heap.Init(&priorityQueue)

	//定义新链表
	newHead := new(ListNode)
	tail := newHead

	for len(priorityQueue) > 0 {
		//优先级队列出队最小的元素
		cur := heap.Pop(&priorityQueue).(*ListNode)
		tail.Next = cur
		tail = tail.Next
		cur = cur.Next
		if cur != nil {
			heap.Push(&priorityQueue, cur)
		}
	}

	return newHead.Next
}
