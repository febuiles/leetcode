package main

import (
	"fmt"
	"math"
)

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	square := 0

	// start from the right since we know what
	// the largest number is. If we knew the smallest
	// we could start from the left?
	for i := right; i >= 0; i-- {
		if math.Abs(float64(nums[left])) < math.Abs(float64(nums[right])) {
			square = nums[right]
			right--
		} else {
			square = nums[left]
			left++
		}
		res[i] = square * square
	}
	return res

}

func main() {

	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
	fmt.Println([]int{0, 1, 9, 16, 100})

}
