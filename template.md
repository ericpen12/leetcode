## 62. 不同路径

### 题目链接：

[https://leetcode-cn.com/problems/unique-paths](https://leetcode-cn.com/problems/unique-paths/)

### 解题思路：

此题和爬楼梯问题类似，如图除了第一列和第一行外，其他任意位置只可能是从上面一格下来，或者从左边一格过来，所以 <u>到某一位置的方式 = 到某位置上一格的方式 +到位置左一格的方式</u>； 一次推导出 dp 方程：

`f[m][n]=f[m-1][n] + f[m][n-1]`; 且 m, n > 1

### 题解代码：

#### 解法一： 二维数组

此时的空间复杂度为 O(m*n)

```go
func uniquePaths(m int, n int) int {
	mn := make([][]int, m)
	for i:= 0; i < m; i++ {
		mn[i] = make([]int, n)
		mn[i][0] = 1
	}
	for j:=0; j < n; j++ {
		mn[0][j] = 1
	}

	for i:=1; i < m; i++ {
		for j:=1; j < n; j++ {
			mn[i][j] = mn[i-1][j] + mn[i][j-1]
		}
	}
	return mn[m-1][n-1]
}
```

[源码](unique_paths.go)



#### 解法二： 一维数组

解法二比解法一稍微难思考一些，解法二是在解法一的思想上，降低空间复杂度， 此时的空间复杂度为 O(n)；

解法一中定义了二维数组，一个作用是用来标记第一列都为 1， 此处，在循环过程中 n 取1， 这样可以保证 n[0]的值不变，并且始终为 1;  mn[j] 表示到上一格的方式， mn[n-1] 表示到左一格的方式。

此处，还可以通过判断 m, n 的大小关系来定义较小的数组

```go
func UniquePaths2(m int, n int) int {
	mn := make([]int, n)
	for i := 0; i < n; i++ {
		mn[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			mn[j] = mn[j] + mn[j-1]
		}
	}
	return mn[n-1]
}
```

[源码](unique_paths.go)



### 代码精选：

```java
public class Solution {
    public int uniquePaths(int m, int n) {
        if(m == 1 || n == 1)
            return 1;
        m--;
        n--;
        if(m < n) {              // Swap, so that m is the bigger number
            m = m + n;
            n = m - n;
            m = m - n;
        }
        long res = 1;
        int j = 1;
        for(int i = m+1; i <= m+n; i++, j++){       // Instead of taking factorial, keep on multiply & divide
            res *= i;
            res /= j;
        }
            
        return (int)res;
    }
}
```

