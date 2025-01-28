package main

import "fmt"

func longestRepeatingCharacterReplacement(s string, k int) int {
	longest := 0
	start := 0
	freq := make(map[byte]int)
	mostFreq := 0

	for end := 0; end < len(s); end++ {
		if _, ok := freq[s[end]]; !ok {
			freq[s[end]] = 1
		} else {
			freq[s[end]]++
		}

		mostFreq = max(mostFreq, freq[s[end]])

		if end-start+1-mostFreq > k {
			freq[s[start]]--
			start++
		}

		longest = max(end-start+1, longest)
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestRepeatingCharacterReplacement("aabccbb", 2))
}
