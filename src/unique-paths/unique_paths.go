package unique_paths

// O(m*n) space
// db: mn[m][n] = mn[m-1][n] + mn[m][n-1]
func UniquePaths(m int, n int) int {
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

// O(n)-space
// dp: mn[n] = mn[n] + mn [n-1]
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

// O(n) space
func UniquePaths3(m int, n int) int {
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

func UniquePaths4(m int, n int) int {
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