package _30_surrounded_regions

var p = []int{0, 0, -1, 1}

// bfs
func solve(board [][]byte) {
	if len(board) < 1 || len(board[0]) < 1 {
		return
	}
	h, w := len(board), len(board[0])

	var queue [][]int
	for i := 0; i < h; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
		}
		if board[i][w-1] == 'O' {
			queue = append(queue, []int{i, w - 1})
		}
	}
	for j := 0; j < w; j++ {
		if board[0][j] == 'O' {
			queue = append(queue, []int{0, j})
		}
		if board[h-1][j] == 'O' {
			queue = append(queue, []int{h - 1, j})
		}
	}

	// bfs
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		x, y := node[0], node[1]
		if board[x][y] == '3' {
			continue
		}
		board[x][y] = '3'

		for i := 0; i < 4; i++ {
			m, n := x+p[i], y+p[3-i]

			if m >= 0 && m < h && n >= 0 && n < w && board[m][n] == 'O' {
				queue = append(queue, []int{m, n})
			}
		}
	}

	for i, v := range board {
		for j, v2 := range v {
			if v2 == '3' {
				board[i][j] = 'O'
			}
			if v2 == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}
