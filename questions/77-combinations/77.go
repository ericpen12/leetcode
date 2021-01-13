package _7_combinations

func combine(n int, k int) [][]int {
	return backtrack(n, k, []int{})
}

func backtrack(n, k int, nums []int) [][]int {
	if len(nums) == k {
		res := make([]int, len(nums))
		copy(res, nums)
		return [][]int{res}
	}
	var result [][]int
	for i := n; i > 0; i-- {
		result = append(result, backtrack(i-1, k, append(nums, i))...)
	}
	return result
}
