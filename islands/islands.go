package main

import "fmt"

func numIslands(grid [][]byte) int {
	numIslands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				if numIslands == 0 || !nearbyIsland(&grid, i, j) {
					numIslands++
				}
			}
		}
	}

	return numIslands
}

func nearbyIsland(grid *[][]byte, i int, j int) bool {
	if i-1 >= 0 {
		if (*grid)[i-1][j] == '1' {
			return true
		}
	}

	if j-1 >= 0 {
		if (*grid)[i][j-1] == '1' {
			return true
		}
	}

	return false
}

func main() {
	input := [][]byte{
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '0', '1'}}

	fmt.Println(numIslands(input))
}
