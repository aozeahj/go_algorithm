package _2_add_two_number

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	cur, cur1, cur2 := newHead, l1, l2
	carry := 0
	for cur1 != nil && cur2 != nil {
		val := cur1.Val + cur2.Val + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val -= 10
		}

		h := new(ListNode)
		h.Val = val

		cur.Next, h.Next = h, cur.Next
		cur = h
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	var l *ListNode

	if cur1 != nil {
		l = cur1
	}
	if cur2 != nil {
		l = cur2
	}

	for l != nil {
		val := l.Val + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val -= 10
		}

		h := new(ListNode)
		h.Val = val

		cur.Next, h.Next = h, cur.Next
		cur = h
		l = l.Next
	}

	if carry == 1 {
		h := new(ListNode)
		h.Val = 1
		cur.Next, h.Next = h, cur.Next
	}

	return newHead.Next
}
