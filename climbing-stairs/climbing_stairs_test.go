package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	target := 8
	n :=  5

	if ClimbStairs(n) != target {
		t.Errorf("ClimbStairs: %d, but target is: %d", ClimbStairs(n), target)
	}
}