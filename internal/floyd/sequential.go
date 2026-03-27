package floyd

import (
	"math"
)

const INF = math.MaxFloat64

func SequentialSP(adjMat [][]float64) [][]float64 {
	n := len(adjMat)
	dist := InitDist(adjMat)

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
	dist, prev := InitDistAndPrev(adjMat)

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
