package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := make([]int, 0)
	row := len(matrix)
	col := len(matrix[0])
	left, right, top, bottom := 0, col-1, 0, row-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		if top < bottom {
			for i := right - 1; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
		}
		if left < right {
			for i := bottom - 1; i >= top+1; i-- {
				res = append(res, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}
