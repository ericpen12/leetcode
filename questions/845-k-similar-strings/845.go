package _45_k_similar_strings

func kSimilarity(A string, B string) int {
	queue := []string{A}
	visited := map[string]bool{}
	var result int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if visited[node] {
				continue
			}
			visited[node] = true

			if node == B {
				return result
			}

			b := []byte(node)
			var i int
			for ; b[i] == B[i]; i++ {
			}
			for j := i + 1; j < len(b); j++ {
				if b[j] == B[j] || b[j] != B[i] {
					continue
				}

				b[i], b[j] = b[j], b
				str := string(b)
				b[i], b[j] = b[j], b[i]

				if str == B {
					return result + 1
				}
				if !visited[str] {
					queue = append(queue, str)
				}
			}

		}
		result++
	}
	return result
}
