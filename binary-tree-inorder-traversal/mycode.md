第一次提交

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    arr := make([]int, 0)
    if root!= nil{
        arr = inorderTraversal(root.Left)
        arr = append(arr, root.Val)
        right := inorderTraversal(root.Right) 
        arr = append(arr, right...)
    }
    return arr
}
```



