package floyd

import (
	"runtime"
	"testing"
)

func TestParallelRowedSP(t *testing.T) {
	fileTestFloyd(t, func(adjMat [][]float64) {
		ParallelRowedSP(adjMat, runtime.NumCPU())
	})
}

func TestParallelRowedSPWithPath(t *testing.T) {
	fileTestFloydWithPath(t, func(adjMat [][]float64, prev [][]int) {
		ParallelRowedSPWithPath(adjMat, prev, runtime.NumCPU())
	})
}

func BenchmarkParallelRowedSP(b *testing.B) {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)
	fileBenchmarkFloyd(b, func(dist [][]float64) {
		ParallelRowedSP(dist, numCpu)
	})
}

func BenchmarkParallelRowedSPRoutines(b *testing.B) {
	benchmarkParallelFloydGoroutines(b, ParallelRowedSP, 5000)
}

func BenchmarkParallelRowedSPProcs(b *testing.B) {
	benchmarkParallelFloydProcs(b, ParallelRowedSP, 5000)
}
