package utils

import "slices"

func GetShortestPath(prev [][]int, u, v int) []int {
	path := []int{}
	for v != u {
		path = append(path, v)
		v = prev[u][v]
		if v == -1 {
			return []int{}
		}
	}
	path = append(path, u)

	slices.Reverse(path)

	return path
}
