package floyd

import (
	"sync"
)

func ParallelRowedSP(dist [][]float64, numOfRoutines int) {
	n := len(dist)

	parallelRowedSP(n, numOfRoutines, func(wgChan <-chan *sync.WaitGroup, startRow, endRow int) {
		floydRowedWorker(wgChan, dist, startRow, endRow)
	})
}

func ParallelRowedSPWithPath(dist [][]float64, prev [][]int, numOfRoutines int) {
	n := len(dist)

	parallelRowedSP(n, numOfRoutines, func(wgChan <-chan *sync.WaitGroup, startRow, endRow int) {
		floydWithPathRowedWorker(wgChan, dist, prev, startRow, endRow)
	})
}

func parallelRowedSP(n, numOfRoutines int, worker func(wgChan <-chan *sync.WaitGroup, startRow, endRow int)) {
	wgChan := make(chan *sync.WaitGroup, numOfRoutines)

	for i := 0; i < numOfRoutines; i++ {
		startRow := (n * i) / numOfRoutines
		endRow := (n * (i + 1)) / numOfRoutines

		go worker(wgChan, startRow, endRow)
	}

	for k := 0; k < n; k++ {
		wg := new(sync.WaitGroup)
		wg.Add(numOfRoutines)

		for i := 0; i < numOfRoutines; i++ {
			wgChan <- wg
		}

		wg.Wait()
	}
}

func floydRowedWorker(wgChan <-chan *sync.WaitGroup, dist [][]float64, startRow, endRow int) {
	n := len(dist)
	for k := 0; k < n; k++ {
		wg := <-wgChan

		floydProcess(dist, k, startRow, endRow, 0, n)

		wg.Done()
		wg.Wait()
	}
}

func floydWithPathRowedWorker(
	wgChan <-chan *sync.WaitGroup,
	dist [][]float64,
	prev [][]int,
	startRow, endRow int,
) {
	n := len(dist)
	for k := 0; k < n; k++ {
		wg := <-wgChan

		floydWithPathProcess(dist, prev, k, startRow, endRow, 0, n)

		wg.Done()
		wg.Wait()
	}
}
