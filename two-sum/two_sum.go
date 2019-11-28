package two_sum

func TwoSum(nums []int, target int) []int {
	sub := make(map[int]int)

	for index, value := range nums {
		if _, isExist := sub[value]; isExist {
			return []int{sub[value], index}
		}
		sub[target - value] = index
	}
	return []int{}
}
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

// 1. 问题类型：求和；
// 2. 解决方法：
//   * 暴力法：遍历数组，挨个加起来判断；时间复杂度 O(n^2)
//   * 求差：可以转化为求差的问题；判断 target 和 当前值的差是否在数组中；时间复杂度 O(n)
// 3. 解题步骤：
//   * 求差：当前值与target 的差
//   * 存差：将差值存入 map，以差为key, index 为值
//   * 判断存在情况：在求差前判断 当前值知否在map中，如果存在，直接返回index，不存在再求差
