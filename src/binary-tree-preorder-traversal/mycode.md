```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    arr := make([]int, 0)
    helper(root, &arr)
    
    return arr
}

func helper(root *TreeNode, arr *[]int) {
    if root != nil {
        *arr = append(*arr, root.Val)
        helper(root.Left, arr)
        helper(root.Right, arr)
    }
}
```

