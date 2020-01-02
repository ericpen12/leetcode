## 94. 二叉树的中序遍历

### 题目链接：

[https://leetcode-cn.com/problems/binary-tree-inorder-traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal)

### 解题思路：

**二叉树中序遍历：**左-根-右

**遍历模板：**

```python
def inorder(self, root):
    if root:
        self.inorder(root.left)
        self.traverse_path.append(root.val)
        self.inorder(root.right)
```



### 题解代码：

#### 解法一：递归

此时的空间复杂度为 O(n)

```go
var traversePath []int

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
```

[源码](binary_tree_inorder_traversal.go)



#### 解法二： 迭代



```go
func inorderTraversal(root *TreeNode) []int {
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
```

[源码](binary_tree_inorder_traversal.go)



### 代码精选：

```java
public List<Integer> inorderTraversal(TreeNode root) {
    List<Integer> list = new ArrayList<Integer>();

    Stack<TreeNode> stack = new Stack<TreeNode>();
    TreeNode cur = root;

    while(cur!=null || !stack.empty()){
        while(cur!=null){
            stack.add(cur);
            cur = cur.left;
        }
        cur = stack.pop();
        list.add(cur.val);
        cur = cur.right;
    }

    return list;
}
```

