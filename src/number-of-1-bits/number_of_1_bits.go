package number_of_1_bits

func HammingWeight(num uint32) int {
	times := 0
	for ; num != 0; times++ {
		num &= num - 1
	}
	return times
}
// 思路：打掉最低位的 1