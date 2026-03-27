package floyd

import (
	"math"
	"sync"
)

func ParallelCachedSPWithPath(adjMat [][]float64, numOfRoutines int) ([][]float64, [][]int) {
	dist, prev := InitDistAndPrev(adjMat)

	blocksPerDim := int(math.Sqrt(float64(numOfRoutines)))

	var wg sync.WaitGroup
	for blockK := range blocksPerDim {
		// phase 1
		floydProcessBlock(dist, prev, blocksPerDim, blockK, blockK, blockK)

		// phase 2
		for blockI := range blocksPerDim {
			if blockI == blockK {
				continue
			}

			wg.Add(2)

			go func(i int) {
				defer wg.Done()
				floydProcessBlock(dist, prev, blocksPerDim, i, blockK, blockK)
			}(blockI)

			go func(j int) {
				defer wg.Done()
				floydProcessBlock(dist, prev, blocksPerDim, blockK, j, blockK)
			}(blockI)
		}

		wg.Wait()

		// phase 3
		for blockI := range blocksPerDim {
			if blockI == blockK {
				continue
			}
			for blockJ := range blocksPerDim {
				if blockJ == blockK {
					continue
				}

				wg.Add(1)
				go func(i, j int) {
					defer wg.Done()
					floydProcessBlock(dist, prev, blocksPerDim, i, j, blockK)
				}(blockI, blockJ)
			}
		}

		wg.Wait()
	}

	return dist, prev
}

func floydProcessBlock(dist [][]float64, prev [][]int, blocksPerDim, blockI, blockJ, blockK int) {
	n := len(dist)
	startI, endI := n*blockI/blocksPerDim, n*(blockI+1)/blocksPerDim
	startJ, endJ := n*blockJ/blocksPerDim, n*(blockJ+1)/blocksPerDim
	startK, endK := n*blockK/blocksPerDim, n*(blockK+1)/blocksPerDim

	for k := startK; k < endK; k++ {
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
}
