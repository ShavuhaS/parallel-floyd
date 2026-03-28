package floyd

import (
	"runtime"
	"testing"
)

func TestParallelBlockedSP(t *testing.T) {
	fileTestFloyd(t, func(adjMat [][]float64) {
		ParallelBlockedSP(adjMat, runtime.NumCPU())
	})
}

func TestParallelBlockedSPWithPath(t *testing.T) {
	fileTestFloydWithPath(t, func(adjMat [][]float64, prev [][]int) {
		ParallelBlockedSPWithPath(adjMat, prev, runtime.NumCPU())
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
