package utils

import "math/rand/v2"

func GenerateMatrix(n, minEdge, maxEdge int, edgeProbability float64) [][]int {
	edgeRange := maxEdge - minEdge
	res := make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				res[i][j] = 0
				continue
			}
			if rand.Float64() < edgeProbability {
				res[i][j] = minEdge + rand.Int()%(edgeRange+1)
			} else {
				res[i][j] = INF
			}
		}
	}

	return res
}
