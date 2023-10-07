package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy, sum := new(ListNode), 0

	for cur := dummy; l1 != nil || l2 != nil || sum != 0; cur = cur.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: sum % 10}
		sum /= 10
	}

	return dummy.Next
}
