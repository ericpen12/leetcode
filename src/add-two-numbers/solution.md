## 2. 两数相加
### 题目链接：

[https://leetcode-cn.com/problems/add-two-numbers/](https://leetcode-cn.com/problems/add-two-numbers/)

### 题解代码：

#### 解法一： 

这个题最大的问题在于**2个不同长度的链表如何处理**，此处给出的处理方法是 **补值**，若当前结点不为空，则取出val, 若为空 val 为 0

```go
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
```

[源码](add_two_numbers.go)



### 代码精选：

```java

```

