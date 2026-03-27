package floyd

import (
	"runtime"
	"testing"
)

func TestParallelPhaseBlockedSP(t *testing.T) {
	fileTestFloyd(t, func(adjMat [][]float64) [][]float64 {
		return ParallelPhaseBlockedSP(adjMat, runtime.NumCPU())
	})
}

func TestParallelPhaseBlockedSPWithPath(t *testing.T) {
	fileTestFloydWithPath(t, func(adjMat [][]float64) ([][]float64, [][]int) {
		return ParallelPhaseBlockedSPWithPath(adjMat, runtime.NumCPU())
	})
}
