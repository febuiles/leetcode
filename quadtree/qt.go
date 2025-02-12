package main

import "fmt"

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func equal(grid *[][]int, l int, x int, y int) bool {
	for i := x; i < x+l; i++ {
		for j := y; j < y+l; j++ {
			// compare against the bottom right boundary value
			if (*grid)[i][j] != (*grid)[x][y] {
				return false
			}
		}
	}

	return true
}

func rec(grid *[][]int, n int, x int, y int) *Node {
	root := Node{}

	if equal(grid, n, x, y) {
		root.IsLeaf = true
	} else {
		root.IsLeaf = false
		root.Val = false

		root.TopLeft = rec(grid, n/2, x, y)
		root.TopRight = rec(grid, n/2, x, y+n/2)
		root.BottomLeft = rec(grid, n/2, x+n/2, y)
		root.BottomRight = rec(grid, n/2, x+n/2, y+n/2)
	}

	if (*grid)[x][y] == 1 {
		root.Val = true
	} else {
		root.Val = false
	}

	return &root
}

func construct(grid [][]int) *Node {
	return rec(&grid, len(grid), 0, 0)
}

func main() {
	input := [][]int{{0, 1}, {1, 2}}
	fmt.Println(construct(input))
}
