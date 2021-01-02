package _00_number_of_islands

func numIslands(grid [][]byte) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	h, w := len(grid), len(grid[0])

	var result int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '0' {
				continue
			}
			bfs(grid, [][]int{{i, j}}, h, w)
			result++
		}
	}
	return result
}

var dr = []int{0, 0, 1, -1}

func bfs(grid [][]byte, queue [][]int, h, w int) {
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		i, j := node[0], node[1]
		if grid[i][j] == '0' {
			continue
		}
		grid[i][j] = '0'

		for k := 0; k < 4; k++ {
			m, n := i+dr[k], j+dr[3-k]
			if m >= 0 && m < h && n >= 0 && n < w && grid[m][n] == '1' {
				queue = append(queue, []int{m, n})
			}
		}
	}
}
