package main

import "fmt"

func bfs(grid *[][]byte, i int, j int) {
	numRows := len((*grid))
	numCols := len((*grid)[0])
	queue := make([][]int, 1)
	queue[0] = []int{i, j}

	for len(queue) > 0 {
		r := queue[0][0]
		c := queue[0][1]
		queue = queue[1:]

		(*grid)[r][c] = '0'

		if r-1 >= 0 && (*grid)[r-1][c] == '1' {
			(*grid)[r-1][c] = '0'
			queue = append(queue, []int{r - 1, c})
		}

		if r+1 < numRows && (*grid)[r+1][c] == '1' {
			(*grid)[r+1][c] = '0'
			queue = append(queue, []int{r + 1, c})
		}

		if c-1 >= 0 && (*grid)[r][c-1] == '1' {
			(*grid)[r][c-1] = '0'
			queue = append(queue, []int{r, c - 1})
		}

		if c+1 < numCols && (*grid)[r][c+1] == '1' {
			(*grid)[r][c+1] = '0'
			queue = append(queue, []int{r, c + 1})
		}
	}
}

func numIslands(grid [][]byte) int {
	numRows := len(grid)
	numCols := len(grid[0])
	numIslands := 0

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if grid[i][j] == '1' {
				numIslands++
				bfs(&grid, i, j)
			}
		}
	}

	return numIslands
}

func main() {

	input := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}

	fmt.Println(numIslands(input))
}
