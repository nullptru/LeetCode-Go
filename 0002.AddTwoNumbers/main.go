package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := l1
	for {
		if l1 == nil {
			return l3
		}
		if l2 != nil {
			l1.Val += l2.Val
			l2 = l2.Next
		}

		if l1.Val >= 10 {
			l1.Val -= 10
			if l1.Next == nil {
				l1.Next = &ListNode{ Val: 0, Next: nil }
			}
			l1.Next.Val++
		}
		if l1.Next == nil && l2 != nil {
			l1.Next = l2
			return l3
		}
		l1 = l1.Next
	}
}