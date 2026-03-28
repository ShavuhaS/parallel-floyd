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

func InitPrev(adjMat [][]float64) [][]int {
	n := len(adjMat)
	prev := make([][]int, n)
	for i := 0; i < n; i++ {
		prev[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if adjMat[i][j] != INF {
				prev[i][j] = i
			} else {
				prev[i][j] = -1
			}
		}
	}
	return prev
}

func floydProcessK(dist [][]float64, startI, endI, startJ, endJ, startK, endK int) {
	for k := startK; k < endK; k++ {
		floydProcess(dist, k, startI, endI, startJ, endJ)
	}
}

func floydProcess(dist [][]float64, k, startI, endI, startJ, endJ int) {
	for i := startI; i < endI; i++ {
		for j := startJ; j < endJ; j++ {
			if dist[i][k] != INF && dist[k][j] != INF {
				newPath := dist[i][k] + dist[k][j]
				if newPath < dist[i][j] {
					dist[i][j] = newPath
				}
			}
		}
	}
}

func floydWithPathProcessK(dist [][]float64, prev [][]int, startI, endI, startJ, endJ, startK, endK int) {
	for k := startK; k < endK; k++ {
		floydWithPathProcess(dist, prev, k, startI, endI, startJ, endJ)
	}
}

func floydWithPathProcess(dist [][]float64, prev [][]int, k, startI, endI, startJ, endJ int) {
	for i := startI; i < endI; i++ {
		for j := startJ; j < endJ; j++ {
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
