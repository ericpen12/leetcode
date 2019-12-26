package add_two_numbers

/**
 * Definition for singly-linked list.
 */


/**
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode pre = new ListNode(0);
        ListNode cur = pre;
        int carry = 0;
        while(l1 != null || l2 != null) {
            int x = l1 == null ? 0 : l1.val;
            int y = l2 == null ? 0 : l2.val;
            int sum = x + y + carry;

            carry = sum / 10;
            sum = sum % 10;
            cur.next = new ListNode(sum);

            cur = cur.next;
            if(l1 != null)
                l1 = l1.next;
            if(l2 != null)
                l2 = l2.next;
        }
        if(carry == 1) {
            cur.next = new ListNode(carry);
        }
        return pre.next;
    }
}
 */
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