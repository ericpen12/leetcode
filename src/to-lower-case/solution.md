## 709. 转换成小写字母

### 题目链接：

[https://leetcode-cn.com/problems/to-lower-case/](https://leetcode-cn.com/problems/to-lower-case/)

### 解题思路：

利用 ascii 码解决问题

### 题解代码：

#### 解法一：位移法

1. 在 ASCII 中， ”a" 是 65， ”A" 是97， 所有字母 ASCII 码间隔 32

2. 将单个字符串进行 **或运算**，结果能实现大写到小写的转换
3. 此处可以换成加法，加32，但是要注意 ASCI I中字母的大小范围

```go
func ToLowerCase(str string) string {
	s := []byte(str)
	for i, v := range s {
		s[i] = v|32
	}
	return string(s)
}
```

[源码](to_lower_case.go)

### 优质代码精选：

```python
class Solution:
    def toLowerCase(self, str): 
        return "".join(chr(ord(c) + 32) if 65 <= ord(c) <= 90 else c for c in str)
```







