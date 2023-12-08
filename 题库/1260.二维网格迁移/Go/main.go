package main

/*
 * @lc app=leetcode.cn id=1260 lang=golang
 *
 * [1260] 二维网格迁移
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) [][]int {
	if len(grid) == 0 {
		return grid
	}

	for i := 0; i < k; i++ {
		move(grid)
	}
	return grid
}

func move(grid [][]int) {
	temp := grid[len(grid)-1][len(grid[0])-1]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j], temp = temp, grid[i][j]
		}
	}
}

// @lc code=end
