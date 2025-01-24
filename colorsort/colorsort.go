package main

import "fmt"

func sortColors(colors []int) []int {
	left := 0
	right := len(colors) - 1
	current := 0

	for current <= right {
		switch colors[current] {
		case 0:
			colors[left], colors[current] = colors[current], colors[left]
			left++
			current++
		case 1:
			current++
		case 2:
			colors[right], colors[current] = colors[current], colors[right]
			right--
		}
	}
	return colors
}

func main() {
	colors := []int{2, 1, 0, 0, 1, 2}
	fmt.Println(sortColors(colors))
}
