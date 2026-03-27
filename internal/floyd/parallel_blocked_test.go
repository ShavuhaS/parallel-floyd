package floyd

import (
	"runtime"
	"testing"
)

func TestParallelBlockedSP(t *testing.T) {
	fileTestFloyd(t, func(adjMat [][]float64) [][]float64 {
		return ParallelBlockedSP(adjMat, runtime.NumCPU())
	})
}

func TestParallelBlockedSPWithPath(t *testing.T) {
	fileTestFloydWithPath(t, func(adjMat [][]float64) ([][]float64, [][]int) {
		return ParallelBlockedSPWithPath(adjMat, runtime.NumCPU())
	})
}
