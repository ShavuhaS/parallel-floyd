package floyd

import (
	"runtime"
	"testing"
)

func TestParallelBlockedSP(t *testing.T) {
	fileTestFloyd(t, func(adjMat [][]float64) {
		ParallelBlockedSP(adjMat, runtime.NumCPU()*runtime.NumCPU())
	})
}

func TestParallelBlockedSPWithPath(t *testing.T) {
	fileTestFloydWithPath(t, func(adjMat [][]float64, prev [][]int) {
		ParallelBlockedSPWithPath(adjMat, prev, runtime.NumCPU()*runtime.NumCPU())
	})
}

func BenchmarkParallelBlockedSP(b *testing.B) {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)
	fileBenchmarkFloyd(b, func(dist [][]float64) {
		ParallelBlockedSP(dist, numCpu)
	})
}

func BenchmarkParallelBlockedSPRoutines(b *testing.B) {
	benchmarkParallelFloydGoroutines(b, ParallelBlockedSP, 5000)
}

func BenchmarkParallelBlockedSPProcs(b *testing.B) {
	benchmarkParallelFloydProcs(b, ParallelBlockedSP, 5000)
}

func BenchmarkParallelBlockedSPBestConfig(b *testing.B) {
	benchmarkParallelFloydConfig(b, ParallelBlockedSP, parallelBlockedRoutineMapping)
}

func parallelBlockedRoutineMapping(v int) int {
	blocksPerDim := v / 600
	blocksPerDim = max(blocksPerDim, 3)
	blocksPerDim = min(blocksPerDim, 8, v)
	numOfRoutines := blocksPerDim * blocksPerDim
	return numOfRoutines
}
