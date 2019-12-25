## 63. 不同路径 II

### 题目链接：

[https://leetcode-cn.com/problems/unique-paths-ii/](https://leetcode-cn.com/problems/unique-paths-ii/)

### 解题思路：

这个问题和 [不同路径](https://leetcode-cn.com/problems/unique-paths/) 的问题类似，此处不同的是，中间有障碍点（可以简单理解为障碍点为 0）；

为了方便处理计算问题，可以新建一个一维数组用来存储计算值，原数组用来判断当前点是否为障碍点。

### 题解代码：

#### 解法：

此时的空间复杂度为 O(m*n)

```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	length := len(obstacleGrid[0])
	mn := make([]int, length)
	mn[0] = 1
	for _, v := range obstacleGrid {
		for j := 0; j < length; j++ {
			if v[j] == 1 {
				mn[j] = 0
				continue
			}
			if j >= 1 {
				mn[j] += mn[j-1]
			}
		}
	}
	return mn[length-1]
}
```

[源码](unique_paths_ii.go)

### 代码精选：

```java

```

