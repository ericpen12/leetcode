package invert_binary_tree

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

// recursion
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree1(root.Left)
	right := invertTree1(root.Right)

	root.Left = right
	root.Right = left

	return root
}
