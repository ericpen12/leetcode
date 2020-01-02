package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var traversePath []int

// recursion
func inorderTraversal(root *TreeNode) []int {
	traversePath = make([]int, 0)
	helper(root)
	return traversePath
}

func helper(root *TreeNode) {
	if root != nil {
		helper(root.Left)
		traversePath = append(traversePath, root.Val)
		helper(root.Right)
	}
}

// Iterative 
func inorderTraversal2(root *TreeNode) []int {
	traversePath := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack)!= 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		traversePath = append(traversePath, cur.Val)
		stack = stack[:len(stack)-1]
		cur = cur.Right

	}

	return traversePath
}