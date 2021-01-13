package _6_permutations

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	return dfs(nums, []int{}, make(map[int]bool, len(nums)))
}

func dfs(nums []int, data []int, visited map[int]bool) [][]int {
	if len(data) == len(nums) {
		res := make([]int, len(data))
		copy(res, data)
		return [][]int{res}
	}
	var result [][]int
	for _, num := range nums {
		if visited[num] {
			continue
		}
		visited[num] = true
		data = append(data, num)
		result = append(result, dfs(nums, data, visited)...)
		// backtrack
		data = data[:len(data)-1]
		visited[num] = false
	}

	return result
}
