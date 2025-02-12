package main

import "fmt"

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	numIslands := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				numIslands++
				dfs(&grid, i, j)
			}
		}
	}

	return numIslands
}

func dfs(grid *[][]byte, i, j int) {
	rows := len(*grid)
	cols := len((*grid)[0])

	(*grid)[i][j] = '0'

	if i-1 >= 0 && (*grid)[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}

	if i+1 < rows && (*grid)[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}

	if j-1 >= 0 && (*grid)[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}

	if j+1 < cols && (*grid)[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
}

func main() {

	input := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}

	fmt.Println(numIslands(input))
}
