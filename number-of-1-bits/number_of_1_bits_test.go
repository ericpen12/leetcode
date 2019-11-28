package number_of_1_bits

import "testing"

func TestHammingWeight(t *testing.T) {
	except := 5
	var nums uint32 = 00000000000000000000000010011011
	result := HammingWeight(nums)
	if result != except {
		t.Errorf("result is: %d, except is:%d", result, except)
	}
}
