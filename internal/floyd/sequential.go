package floyd

import "math"

const INF = math.MaxInt

func SequentialSP(adjMat [][]int) [][]int {
	n := len(adjMat)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if adjMat[i][k] != INF && adjMat[k][j] != INF {
					newPath := adjMat[i][k] + adjMat[k][j]
					if newPath < adjMat[i][j] {
						adjMat[i][j] = newPath
					}
				}
			}
		}
	}
	return adjMat
}
