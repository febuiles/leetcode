package main

import (
	"fmt"
)

func findLongestSubstring(str string) int {
	if len(str) == 0 {
		return 0
	}

	start, longest := 0, 0
	dictionary := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		if prevIndex, ok := dictionary[str[i]]; ok && prevIndex >= start {
			start = prevIndex + 1
		}

		dictionary[str[i]] = i

		if i-start+1 > longest {
			longest = i - start + 1
		}
	}

	return longest
}

func main() {
	input := "pwkeww"

	fmt.Println(findLongestSubstring(input))
}
