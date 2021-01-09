package _6_pond_sizes_lcci

import "sort"

func pondSizes(land [][]int) []int {
	if len(land) == 0 || len(land[0]) == 0 {
		return nil
	}
	m, n := len(land), len(land[0])
	var result []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if size := dfs(land, i, j, m, n); size > 0 {
				result = append(result, size)
			}
		}
	}
	sort.Ints(result)
	return result
}

var dr = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func dfs(land [][]int, i, j, m, n int) int {
	if i < 0 || i >= m || j < 0 || j >= n || land[i][j] != 0 {
		return 0
	}
	land[i][j] = 1
	var result = 1
	for _, v := range dr {
		result += dfs(land, i+v[0], j+v[1], m, n)
	}
	return result
}
