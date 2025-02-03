package main

import "fmt"

func longestLoop(s string) string {

	coords := make(map[[2]int]int)
	coords[[2]int{0, 0}] = -1

	x, y := 0, 0
	longest := 0
	res := ""
	for i, ch := range s {

		switch ch {
		case 'u':
			y++
		case 'd':
			y--
		case 'l':
			x--
		case 'r':
			x++
		}
		if prevIndex, exists := coords[[2]int{x, y}]; exists {
			length := i - prevIndex
			if length > longest {
				longest = length
				res = s[prevIndex+1 : i+1]
			}
		}
		coords[[2]int{x, y}] = i
	}

	return res
}

func main() {
	fmt.Println(longestLoop("dlur"))
}
