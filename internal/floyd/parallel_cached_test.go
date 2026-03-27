package floyd

import (
	"runtime"
	"testing"
)

func TestParallelCachedSPWithPath(t *testing.T) {
	fileTestFloydWithPath(t, func(adjMat [][]float64) ([][]float64, [][]int) {
		return ParallelCachedSPWithPath(adjMat, runtime.NumCPU())
	})
}
