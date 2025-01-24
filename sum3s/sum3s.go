package main

import (
	"fmt"
	"sort"
)

func findSumOfThree(nums []int, target int) bool {

	sort.Sort(sort.IntSlice(nums))
	low, high, sum := 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {
		low = i + 1
		high = len(nums) - 1

		for low < high {
			sum = nums[i] + nums[low] + nums[high]

			if sum == target {
				fmt.Println(nums[i], nums[low], nums[high])
				return true
			} else if sum < target {
				low += 1
			} else {
				high -= 1
			}

		}
	}
	return false
}

func main() {
	//	1 2 3 4 5 7 8

	test := []int{3, 7, 1, 2, 8, 4, 5}
	findSumOfThree(test, 15)

}
