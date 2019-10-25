package main

import "fmt"

func generateParenthesis(n int) []string {
	arr := make([]string, 0)
	generate(0, 0, n,"", &arr)
	return arr
}

func generate(left int, right int, n int, str string, i *[]string) {
	if left == n && right == n {
		*i = append(*i, str)
		return
	}
	 if left < n {
		 generate(left + 1, right, n, str + "(", i)
	 }
	if right < left {
		generate(left, right + 1, n, str + ")", i)
	}
}


func main()  {
	n := 3
	arr := generateParenthesis(n)
	fmt.Println(arr)
}