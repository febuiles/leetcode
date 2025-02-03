package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	files := make(map[string][]string)
	for _, path := range paths {
		parts := strings.Split(path, " ")
		dir, rest := parts[0], parts[1]
		parts = strings.Split(rest, "(")
		filename := parts[0]
		content := parts[1]
		path = dir + "/" + filename

		files[content] = append(files[content], path)
	}

	res := [][]string{}
	for _, files := range files {
		if len(files) > 1 {
			res = append(res, files)
		}
	}

	return res

}

func main() {
	input := []string{"root/app1 1.log(error1)", "root/app2 2.log(error2)", "root/app2 3.log(error1)", "root/app4 5.log(error2)"}
	fmt.Println(findDuplicate(input))
}
