package floyd

import (
	"runtime"
	"testing"
)

func TestParallelRowedSP(t *testing.T) {
	fileTestFloyd(t, func(adjMat [][]float64) [][]float64 {
		return ParallelRowedSP(adjMat, runtime.NumCPU())
	})
}

func TestParallelRowedSPWithPath(t *testing.T) {
	fileTestFloydWithPath(t, func(adjMat [][]float64) ([][]float64, [][]int) {
		return ParallelRowedSPWithPath(adjMat, runtime.NumCPU())
	})
}
