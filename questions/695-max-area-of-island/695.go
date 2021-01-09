package _95_max_area_of_island

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var result int
	m, n := len(grid), len(grid[0])
	var visited = map[int]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 && !visited[i*n+j] {
				num := bfs(grid, visited, [][]int{{i, j}}, m, n)
				if num > result {
					result = num
				}
			}
		}
	}
	return result
}

var dr = []int{0, 0, -1, 1}

func bfs(grid [][]int, visited map[int]bool, stack [][]int, m, n int) int {
	var result int
	for len(stack) > 0 {
		index := len(stack) - 1
		node := stack[index]
		stack = stack[:index]
		i, j := node[0], node[1]
		if visited[i*n+j] {
			continue
		}
		visited[i*n+j] = true
		result++
		for k := 0; k < 4; k++ {
			x, y := i+dr[k], j+dr[3-k]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
				stack = append(stack, []int{x, y})
			}
		}
	}
	return result
}
