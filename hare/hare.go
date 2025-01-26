package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

func main() {
	input := []int{1, 6, 3, 2, 3, 5, 4}
	fmt.Println(findDuplicate(input))
}
