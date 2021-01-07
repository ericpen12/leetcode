package _01_symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}

			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		for i := 0; i < len(queue)/2; i++ {
			left, right := queue[i], queue[len(queue)-1-i]
			if (left == right) || (left != nil && right != nil && left.Val == right.Val) {
				continue
			}
			return false
		}
	}
	return true
}
