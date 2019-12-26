package add_two_numbers

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{
		Val:  0,
		Next: nil,
	}

	cur := pre
	carry := 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry

		carry = sum / 10
		sum = sum % 10

		cur.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		cur = cur.Next
	}
	if carry == 1 {
		cur.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return pre.Next
}

// 1. 分析：链表问题，需要遍历链表
// 2. 思路一：使用递归，遍历到最底层，然后在原节点开始加法运算，返回是否进一