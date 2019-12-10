package unique_paths

func UniquePaths(m int, n int) int {
	less, more := m, n
	if m > n {
		less =n
		more = m
	}
	
	mn := make([]int, less)
	for i, _ := range mn {
		mn[i] = 1
	}
	for i:=1; i < more; i++ {
		for j:=1; j < less; j++ {
			mn[j] = mn[j-1]+mn[j]
		}
	}
	return mn[less-1]
}

// 1. 动态规划问题
// 2. dp 方程：f(i,j) = f(i, j-1) + f(i-1,j);