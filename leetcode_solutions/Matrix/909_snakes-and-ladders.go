# Snakes and Ladders
# Difficulty: Medium
# Language: golang
# Topic: Matrix
# Tags: Array, Breadth-First Search, Matrix
# Link: https://leetcode.com/problems/snakes-and-ladders/

func snakesAndLadders(board [][]int) int {
	n := len(board)
	if board[0][0] != -1 || n == 1 {
		return -1
	}
	answer := 400

	visited := map[int]int{}
	for i := 1; i <= n*n; i++ {
		visited[i] = -1
	}
	visited[1] = 0

	q := []int{}
	q = append(q, 1)

	for len(q) > 0 {
		k := q[0]
		q = q[1:]

		if k == n*n {
			//	if answer > visited[k] {
			//		answer = visited[k]
			//	}
			return visited[k]
		}

		for d := 1; d <= 6 && k+d <= n*n; d++ {
			i, j := numberToCoordinate(n, k+d)[0], numberToCoordinate(n, k+d)[1]
			if board[i][j] != -1 {
				if visited[board[i][j]] == -1 || visited[board[i][j]] > visited[k]+1 {
					visited[board[i][j]] = visited[k] + 1
					q = append(q, board[i][j])
					
				}
                continue
			}

			//if board[i][j] == -1 &&  visited[k+d] == -1 {
			if (visited[k+d] == -1 || visited[k+d] > visited[k]+1) {
				visited[k+d] = visited[k] + 1
				q = append(q, k+d)

			}
		}
	}

	if answer == 400 {
		return -1
	}

	return answer
}

func numberToCoordinate(size, k int) []int {
	k = k - 1
	row := k / size
	column := (k % size)
	if row%2 == 1 {
		column = size - column - 1
	}
	row = size - row - 1
	return []int{row, column}
}

