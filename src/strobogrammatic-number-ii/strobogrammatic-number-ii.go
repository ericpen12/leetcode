package main

import (
	"fmt"
)

func findStrobogrammatic(n int) []string {
	arr := make([]string, 0)
	helper(1, n, "", &arr)
	return arr
}

func helper(level, n int, str string, arr *[]string) {
	nums := map[int]string{
		0: "0",
		1: "1",
		6: "9",
		8: "8",
		9: "6",
	}

	// terminator
	if level > 3 {
		// process result
		//*arr = append(*arr, str)
		fmt.Println(str)
		return
	}

	// process current logic

	for key, val := range nums {
		if key == 0 && level == 1 {
			continue
		}

		if left < n/2 {
			helper(level + 1, str + val, arr)
		}
	}

	// drill down

	// restore current status
}

func main() {
	fmt.Println(findStrobogrammatic(3))
}
