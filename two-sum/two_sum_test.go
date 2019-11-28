package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := TwoSum(nums, target)
	right := []int{0, 1}
	for i, v := range right {
		if res[i] != v {
			t.Errorf("result: %d, right: %d", res[i], v)
		}
	}

	//因为 nums[0] + nums[1] = 2 + 7 = 9
	//所以返回 [0, 1]
}
