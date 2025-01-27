package main

import (
	"fmt"
)

func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

func sumOfSquaredDigits(number int) int {
	totalSum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		totalSum += pow(digit, 2)
	}
	return totalSum
}

func isHappy(num int) bool {
	slowPointer := num
	fastPointer := sumOfSquaredDigits(num)

	for fastPointer != 1 && slowPointer != fastPointer {
		slowPointer = sumOfSquaredDigits(slowPointer)
		fastPointer = sumOfSquaredDigits(sumOfSquaredDigits(fastPointer))
	}

	if fastPointer == 1 {
		return true
	}
	return false
}
func main() {
	input := 2147483646
	fmt.Println(isHappy(input))
}
