## 144. 二叉树的前序遍历

### 题目链接：

[https://leetcode-cn.com/problems/binary-tree-preorder-traversal/](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)

### 解题思路：

前序遍历：根-左-右


### 题解代码：

#### 解法一： 递归

使用递归更加方便思考，使用递归的基本模板就好了

```go
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
```

[源码](binary_tree_preorder_traversal.go)

#### 解法二： 迭代

迭代的思路比较难理解，通过手动维护一个栈，想法类似深度优先搜索

```go
func preorderTraversal(root *TreeNode) []int {
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
```

[源码](binary_tree_preorder_traversal.go)



### 代码精选：

```java
class Solution {
    public List<Integer> preorderTraversal(TreeNode root) {
        Deque<TreeNode> stack = new LinkedList<>();
        TreeNode p = root;
        List<Integer> res = new ArrayList<>();
        while (p != null || !stack.isEmpty()) {
            while (p != null) {
                res.add(p.val);
                stack.push(p);
                p = p.left;
            }
            p = stack.pop().right;
        }
        return res;
    }
}
```

