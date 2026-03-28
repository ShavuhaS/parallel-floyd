package floyd

import (
	"math"
	"sync"
)

func ParallelBlockedSP(dist [][]float64, numOfRoutines int) {
	n := len(dist)

	parallelBlockedSP(n, numOfRoutines, func(wgChan <-chan *sync.WaitGroup, startI, endI, startJ, endJ int) {
		floydBlockedWorker(wgChan, dist, startI, endI, startJ, endJ)
	})
}

func ParallelBlockedSPWithPath(dist [][]float64, prev [][]int, numOfRoutines int) {
	n := len(dist)

	parallelBlockedSP(n, numOfRoutines, func(wgChan <-chan *sync.WaitGroup, startI, endI, startJ, endJ int) {
		floydWithPathBlockedWorker(wgChan, dist, prev, startI, endI, startJ, endJ)
	})
}

func parallelBlockedSP(
	n,
	numOfRoutines int,
	worker func(wgChan <-chan *sync.WaitGroup, startI, endI, startJ, endJ int),
) {
	wgChan := make(chan *sync.WaitGroup, numOfRoutines)

	blocksPerDim := int(math.Sqrt(float64(numOfRoutines)))
	for i := 0; i < blocksPerDim; i++ {
		for j := 0; j < blocksPerDim; j++ {
			startRow := (n * i) / blocksPerDim
			endRow := (n * (i + 1)) / blocksPerDim
			startCol := (n * j) / blocksPerDim
			endCol := (n * (j + 1)) / blocksPerDim

			go worker(wgChan, startRow, endRow, startCol, endCol)
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
}

func floydBlockedWorker(
	wgChan <-chan *sync.WaitGroup,
	dist [][]float64,
	startI, endI int,
	startJ, endJ int,
) {
	n := len(dist)
	for k := 0; k < n; k++ {
		wg := <-wgChan

		floydProcess(dist, k, startI, endI, startJ, endJ)

		wg.Done()
		wg.Wait()
	}
}

func floydWithPathBlockedWorker(
	wgChan <-chan *sync.WaitGroup,
	dist [][]float64,
	prev [][]int,
	startI, endI int,
	startJ, endJ int,
) {
	n := len(dist)
	for k := 0; k < n; k++ {
		wg := <-wgChan

		floydWithPathProcess(dist, prev, k, startI, endI, startJ, endJ)

		wg.Done()
		wg.Wait()
	}
}
