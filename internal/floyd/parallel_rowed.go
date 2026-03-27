package floyd

import (
	"sync"
)

func ParallelRowedSP(adjMat [][]float64, numOfRoutines int) [][]float64 {
	n := len(adjMat)
	dist := InitDist(adjMat)

	wgChan := make(chan *sync.WaitGroup, numOfRoutines)

	for i := 0; i < numOfRoutines; i++ {
		startRow := (n * i) / numOfRoutines
		endRow := (n * (i + 1)) / numOfRoutines

		go floydRowedWorker(wgChan, dist, n, startRow, endRow)
	}

	for k := 0; k < n; k++ {
		wg := new(sync.WaitGroup)
		wg.Add(numOfRoutines)

		for i := 0; i < numOfRoutines; i++ {
			wgChan <- wg
		}

		wg.Wait()
	}

	return dist
}

func ParallelRowedSPWithPath(adjMat [][]float64, numOfRoutines int) ([][]float64, [][]int) {
	n := len(adjMat)
	dist, prev := InitDistAndPrev(adjMat)

	wgChan := make(chan *sync.WaitGroup, numOfRoutines)

	for i := 0; i < numOfRoutines; i++ {
		startRow := (n * i) / numOfRoutines
		endRow := (n * (i + 1)) / numOfRoutines

		go floydWithPathRowedWorker(wgChan, dist, prev, n, startRow, endRow)
	}

	for k := 0; k < n; k++ {
		wg := new(sync.WaitGroup)
		wg.Add(numOfRoutines)

		for i := 0; i < numOfRoutines; i++ {
			wgChan <- wg
		}

		wg.Wait()
	}

	return dist, prev
}

func floydRowedWorker(wgChan <-chan *sync.WaitGroup, dist [][]float64, n, startRow, endRow int) {
	for k := 0; k < n; k++ {
		wg := <-wgChan

		for i := startRow; i < endRow; i++ {
			for j := 0; j < n; j++ {
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

func floydWithPathRowedWorker(
	wgChan <-chan *sync.WaitGroup,
	dist [][]float64,
	prev [][]int,
	n, startRow, endRow int,
) {
	for k := 0; k < n; k++ {
		wg := <-wgChan

		for i := startRow; i < endRow; i++ {
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

		wg.Done()
		wg.Wait()
	}
}
