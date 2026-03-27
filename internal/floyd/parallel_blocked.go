package floyd

import (
	"math"
	"sync"
)

func ParallelBlockedSP(adjMat [][]float64, numOfRoutines int) [][]float64 {
	n := len(adjMat)
	dist := InitDist(adjMat)

	wgChan := make(chan *sync.WaitGroup, numOfRoutines)

	blocksPerDim := int(math.Sqrt(float64(numOfRoutines)))
	for i := 0; i < blocksPerDim; i++ {
		for j := 0; j < blocksPerDim; j++ {
			startRow := (n * i) / blocksPerDim
			endRow := (n * (i + 1)) / blocksPerDim
			startCol := (n * j) / blocksPerDim
			endCol := (n * (j + 1)) / blocksPerDim

			go floydBlockedWorker(wgChan, dist, n, startRow, endRow, startCol, endCol)
		}
	}

	actualRoutineCount := blocksPerDim * blocksPerDim
	for k := 0; k < n; k++ {
		wg := new(sync.WaitGroup)
		wg.Add(actualRoutineCount)

		for i := 0; i < actualRoutineCount; i++ {
			wgChan <- wg
		}

		wg.Wait()
	}

	return dist
}

func ParallelBlockedSPWithPath(adjMat [][]float64, numOfRoutines int) ([][]float64, [][]int) {
	n := len(adjMat)
	dist, prev := InitDistAndPrev(adjMat)

	wgChan := make(chan *sync.WaitGroup, numOfRoutines)

	blocksPerDim := int(math.Sqrt(float64(numOfRoutines)))
	for i := 0; i < blocksPerDim; i++ {
		for j := 0; j < blocksPerDim; j++ {
			startRow := (n * i) / blocksPerDim
			endRow := (n * (i + 1)) / blocksPerDim
			startCol := (n * j) / blocksPerDim
			endCol := (n * (j + 1)) / blocksPerDim

			go floydWithPathBlockedWorker(wgChan, dist, prev, n, startRow, endRow, startCol, endCol)
		}
	}

	actualRoutineCount := blocksPerDim * blocksPerDim
	for k := 0; k < n; k++ {
		wg := new(sync.WaitGroup)
		wg.Add(actualRoutineCount)

		for i := 0; i < actualRoutineCount; i++ {
			wgChan <- wg
		}

		wg.Wait()
	}

	return dist, prev
}

func floydBlockedWorker(
	wgChan <-chan *sync.WaitGroup,
	dist [][]float64,
	n,
	startRow, endRow,
	startCol, endCol int,
) {
	for k := 0; k < n; k++ {
		wg := <-wgChan

		for i := startRow; i < endRow; i++ {
			for j := startCol; j < endCol; j++ {
				if dist[i][k] != INF && dist[k][j] != INF {
					newPath := dist[i][k] + dist[k][j]
					if newPath < dist[i][j] {
						dist[i][j] = newPath
					}
				}
			}
		}

		wg.Done()
		wg.Wait()
	}
}

func floydWithPathBlockedWorker(
	wgChan <-chan *sync.WaitGroup,
	dist [][]float64,
	prev [][]int,
	n,
	startRow, endRow,
	startCol, endCol int,
) {
	for k := 0; k < n; k++ {
		wg := <-wgChan

		for i := startRow; i < endRow; i++ {
			for j := startCol; j < endCol; j++ {
				if dist[i][k] != INF && dist[k][j] != INF {
					newPath := dist[i][k] + dist[k][j]
					if newPath < dist[i][j] {
						dist[i][j] = newPath
						prev[i][j] = prev[k][j]
					}
				}
			}
		}

		wg.Done()
		wg.Wait()
	}
}
