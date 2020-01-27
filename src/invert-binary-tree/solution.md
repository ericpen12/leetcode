# 226. 翻转二叉树

## 题目链接：

[https://leetcode-cn.com/problems/invert-binary-tree/](https://leetcode-cn.com/problems/invert-binary-tree/)

## 解法一：递归法

此题使用递归解题比较容易理解， 每次逐渐下探，并返回该节点前，同时将左右结点互换

```go
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}

```

[源码](invert_binary_tree.go)

