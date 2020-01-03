package binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion
var traversePath []int
func preorderTraversal(root *TreeNode) []int {
	traversePath = make([]int, 0)
	help(root)
	return traversePath
}

func help(root *TreeNode) {
	if root != nil {
		traversePath = append(traversePath, root.Val)
		help(root.Left)
		help(root.Right)
	}
}


// Iterative
func preorderTraversal2(root *TreeNode) []int {
	arr := make([]int, 0)
	stack := make([]*TreeNode, 0)
	curr := root

	for curr != nil || len(stack) != 0 {
		for curr != nil {
			arr = append(arr, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return arr
}
