package floyd

import "slices"

func InitDist(adjMat [][]float64) [][]float64 {
	n := len(adjMat)
	dist := make([][]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = slices.Clone(adjMat[i])
	}
	return dist
}

func InitDistAndPrev(adjMat [][]float64) ([][]float64, [][]int) {
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
	return dist, prev
}
