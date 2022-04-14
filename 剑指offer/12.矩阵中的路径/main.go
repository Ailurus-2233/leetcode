package main

func exist(board [][]byte, word string) bool {
	ans := false

outlabel:
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if ans {
				break outlabel
			}
			if board[x][y] == word[0] {
				ans = isExist(board, word, x, y)
			}
		}
	}
	return ans
}

func isExist(board [][]byte, word string, x int, y int) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}
	ans := board[x][y] == word[0]
	if ans && len(word) != 1 {
		temp := board[x][y]
		board[x][y] = ' '
		ans = ans && (isExist(board, word[1:], x-1, y) || isExist(board, word[1:], x+1, y) || isExist(board, word[1:], x, y-1) || isExist(board, word[1:], x, y+1))
		board[x][y] = temp
	}
	return ans
}
