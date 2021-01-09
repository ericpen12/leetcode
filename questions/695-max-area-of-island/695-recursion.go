package _95_max_area_of_island

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var result int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if num := dfs(grid, i, j, m, n); num > result {
				result = num
			}
		}
	}
	return result
}

func dfs(grid [][]int, i, j, m, n int) int {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return dfs(grid, i+1, j, m, n) +
		dfs(grid, i-1, j, m, n) +
		dfs(grid, i, j+1, m, n) +
		dfs(grid, i, j-1, m, n) + 1
}
