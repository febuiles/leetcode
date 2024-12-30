package main

import "fmt"

func convert(s string, numRows int) string {
	mat := make([][]rune, numRows)
	for i := range mat {
		mat[i] = []rune{}
	}

	idx, d := 0, 1
	for _, r := range s {
		mat[idx] = append(mat[idx], r)
		if idx == 0 {
			d = 1
		} else if idx == numRows-1 {
			d = -1
		}

		idx += d
	}

	var result []rune
	for _, row := range mat {
		result = append(result, row...)
	}

	return string(result)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))

}
