package generate_parentheses

var arr []string

func generateParenthesis(n int) []string {
	arr := make([]string, 0)
	generate(0, 0, n,"")
	return arr
}

func generate(left int, right int, n int, str string) {
	if left == n && right == n {
		arr = append(arr, str)
		return
	}
	 if left < n {
		 generate(left + 1, right, n, str + "(")
	 }
	if right < left {
		generate(left, right + 1, n, str + ")")
	}
}