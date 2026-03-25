package floyd

import (
	"math"
	"slices"
)

const INF = math.MaxFloat64

func SequentialSP(adjMat [][]float64) [][]float64 {
	n := len(adjMat)
	dist := make([][]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = slices.Clone(adjMat[i])
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != INF && dist[k][j] != INF {
					newPath := dist[i][k] + dist[k][j]
					if newPath < dist[i][j] {
						dist[i][j] = newPath
					}
				}
			}
		}
	}

	return dist
}

func SequentialSPWithPath(adjMat [][]float64) ([][]float64, [][]int) {
	n := len(adjMat)
	dist := make([][]float64, n)
	prev := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = slices.Clone(adjMat[i])
		prev[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				prev[i][i] = i
				continue
			}
			if dist[i][j] != INF {
				prev[i][j] = i
			} else {
				prev[i][j] = -1
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != INF && dist[k][j] != INF {
					newPath := dist[i][k] + dist[k][j]
					if newPath < dist[i][j] {
						dist[i][j] = newPath
						prev[i][j] = prev[k][j]
					}
				}
			}
		}
	}

	return dist, prev
}
