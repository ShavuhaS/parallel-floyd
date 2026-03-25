package utils

import (
	"math"
	"math/rand/v2"
)

func GenerateMatrix(n int, minEdge, maxEdge, edgeProbability float64) [][]float64 {
	edgeRange := maxEdge - minEdge
	res := make([][]float64, n)

	for i := 0; i < n; i++ {
		res[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				res[i][j] = 0
				continue
			}
			if rand.Float64() < edgeProbability {
				res[i][j] = math.Ceil(minEdge + rand.Float64()*edgeRange)
			} else {
				res[i][j] = INF
			}
		}
	}

	return res
}
