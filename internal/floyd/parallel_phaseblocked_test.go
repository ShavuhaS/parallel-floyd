package floyd

import (
	"math"
	"runtime"
	"testing"
)

func TestParallelPhaseBlockedSP(t *testing.T) {
	fileTestFloyd(t, func(adjMat [][]float64) {
		ParallelPhaseBlockedSP(adjMat, runtime.NumCPU()*runtime.NumCPU())
	})
}

func TestParallelPhaseBlockedSPWithPath(t *testing.T) {
	fileTestFloydWithPath(t, func(adjMat [][]float64, prev [][]int) {
		ParallelPhaseBlockedSPWithPath(adjMat, prev, runtime.NumCPU()*runtime.NumCPU())
	})
}

func BenchmarkParallelPhaseBlockedSP(b *testing.B) {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)
	fileBenchmarkFloyd(b, func(dist [][]float64) {
		ParallelPhaseBlockedSP(dist, numCpu)
	})
}

func BenchmarkParallelPhaseBlockedSPRoutines(b *testing.B) {
	benchmarkParallelFloydGoroutines(b, ParallelPhaseBlockedSP, 5000)
}

func BenchmarkParallelPhaseBlockedSPProcs(b *testing.B) {
	benchmarkParallelFloydProcs(b, ParallelPhaseBlockedSP, 5000)
}

func BenchmarkParallelPhaseBlockedSPBestConfig(b *testing.B) {
	benchmarkParallelFloydConfig(b, ParallelPhaseBlockedSP, parallelPhaseBlockedRoutineMapping)
}

func parallelPhaseBlockedRoutineMapping(v int) int {
	blocksPerDim := int(math.Ceil(float64(v) / 150.0))
	blocksPerDim = max(blocksPerDim, 3)
	blocksPerDim = min(blocksPerDim, v)
	numOfRoutines := blocksPerDim * blocksPerDim
	return numOfRoutines
}
