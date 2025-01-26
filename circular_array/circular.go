package main

import (
	"fmt"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func nextStep(pointer int, value int, size int) int {
	result := (pointer + value) % size
	if result < 0 {
		result += size
	}
	return result
}

func isNotCycle(nums []int, prevDirection bool, pointer int) bool {
	currDirection := nums[pointer] >= 0

	if (prevDirection != currDirection) || (abs(nums[pointer])%len(nums) == 0) {
		return true
	} else {
		return false
	}
}

func circularArrayLoop(nums []int) bool {
	size := len(nums)

	for i := 0; i < size; i++ {

		slow, fast := i, i

		forward := nums[i] > 0

		for {
			slow = nextStep(slow, nums[slow], size)
			if isNotCycle(nums, forward, slow) {
				break
			}

			fast = nextStep(fast, nums[fast], size)
			if isNotCycle(nums, forward, fast) {
				break
			}

			fast = nextStep(fast, nums[fast], size)
			if isNotCycle(nums, forward, fast) {
				break
			}

			if slow == fast {
				return true
			}
		}
	}

	return false
}

func main() {

	input := []int{-2, -3, -9}
	fmt.Println(circularArrayLoop(input))
}
