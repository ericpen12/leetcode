## 22. 括号生成

### 题目链接：

[https://leetcode-cn.com/problems/generate-parentheses](https://leetcode-cn.com/problems/generate-parentheses)

### 题解代码：

#### 解法一： 递归

**思路**：

1. 利用递归给字符串逐渐递加括号
2. 每次增加括号前判断是否能添加对应的左右括号
3. 判断条件：添加左括号---左括号数量小于 n, 添加右括号---右括号数量小于左括号

**时间复杂度**：O(n)

**空间复杂度**：

```go
var arr []string

func generateParenthesis(n int) []string {
	arr = make([]string, 0)
	generate(0, 0, n, "")
	return arr
}

func generate(left int, right int, n int, str string) {
	if len(str) == 2*n {
		arr = append(arr, str)
		return
	}
	 if left < n {
		 generate(left + 1, right, n, str + "(")
	 }
	if right < left {
		generate(left, right + 1, n, str + ")")
	}
}
```

[源码](generate_parentheses.go)


