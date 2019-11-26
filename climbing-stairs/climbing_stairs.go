package climbing_stairs

var hashTable = make(map[int]int)

func ClimbStairs(n int) int {
	if n < 3 {
		return n
	}
	if _, isExist := hashTable[n]; !isExist {
		hashTable[n] = ClimbStairs(n - 1) + ClimbStairs(n - 2)
	}
	return hashTable[n]
}

// 分析：
// 问题：典型的斐波那契问题 f(n)=f(n-1)+f(n-2)
// 求解方式：递归，for loop
// 递归：
// 	  终止条件: n < 3
// 	  递归公式：f(n)=f(n-1)+f(n-2); 带入 n-1; n-2
//    返回值：return f(n)
//    时间复杂度：若不经过处理，O(2^n); 可以使用 hashTable 记录已经处理过的数据；此时 O(n)