package _45_k_similar_strings

func kSimilarity2(A string, B string) int {
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
			var ii int
			for ; b[ii] == B[ii]; ii++ {
			}
			for j := ii + 1; j < len(b); j++ {
				temp := []byte(node)
				if temp[j] == B[j] || temp[j] != B[ii] {
					continue
				}

				temp[ii], temp[j] = temp[j], temp[ii]

				if string(temp) == B {
					return result + 1
				}

				if !visited[string(temp)] {
					queue = append(queue, string(temp))
				}
			}

		}

		result++
	}
	return result
}
