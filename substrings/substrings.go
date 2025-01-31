package main

import (
	"fmt"
	"math"
)

func findRepeatedSequences(dna string, k int) Set {
	length := len(dna)
	mapping := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}
	baseValue := 4

	if length < k {
		return *NewSet()
	}

	numbers := make([]int, length)
	for i, ch := range dna {
		numbers[i] = mapping[ch]
	}

	hashValue := 0

	hashSet := *NewSet()
	output := *NewSet()
	// use a rolling hash for each number
	// computing the hash takes O(k) initially, O(1) for updates
	// lookup is O(1)
	// if we were using substrings, checking against the set would take O(nk)
	for i := 0; i < length-k+1; i++ {
		// calculating the first hash is independent of other vals
		if i == 0 {
			for j := i; j < k; j++ {

				hashValue += numbers[j] * int(math.Pow(float64(baseValue), float64(k-j-1)))
			}
		} else {
			previousHashValue := hashValue
			hashValue = ((previousHashValue - numbers[i-1]*int(math.Pow(float64(baseValue), float64(k-1)))) * baseValue) + numbers[i+k-1]

		}

		if hashSet.Exists(hashValue) {
			output.Add(dna[i : i+k])
		}
		hashSet.Add(hashValue)
	}

	return output
}

func main() {
	str := "GACGACGAC"
	k := 3
	fmt.Println(findRepeatedSequences(str, k))
}

type Set struct {
	hashMap map[interface{}]bool
}

// NewSet will initialize and return a new object of Set.
func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[interface{}]bool)
	return s
}

// Add will add the value in the Set.
func (s *Set) Add(value interface{}) {
	s.hashMap[value] = true
}

// Delete will delete the value from the set.
func (s *Set) Delete(value interface{}) {
	delete(s.hashMap, value)
}

// Exists will check if the value exists in the set or not.
func (s *Set) Exists(value interface{}) bool {
	_, ok := s.hashMap[value]
	return ok
}
