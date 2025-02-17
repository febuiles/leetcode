package main

import (
	"math"
)

func networkDelayTime(times [][]int, n, k int) int {
	nodes := make(map[int][][]int)
	visited := make(map[int]bool)
	res := 0

	for _, data := range times {
		src := data[0]
		dst := data[1]
		w := data[2]

		nodes[src] = append(nodes[src], []int{dst, w})
	}

	pq := make([][]int, 0)
	pq = append(pq, []int{0, k})

	for len(pq) > 0 {
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i][0] < pq[minIdx][0] {
				minIdx = i
			}
		}

		cur := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)
		time := cur[0]
		node := cur[1]

		if visited[node] {
			continue
		}

		visited[node] = true
		res = int(math.Max(float64(time), float64(res)))
		neighbors := nodes[node]
		for _, neighbor := range neighbors {
			n := neighbor[0]
			w := neighbor[1]
			if !visited[n] {
				newTime := time + w
				pq = append(pq, []int{newTime, n})
			}
		}

	}

	if len(visited) != n {
		return -1
	} else {
		return res
	}

}
