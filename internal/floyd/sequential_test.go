package floyd

import (
	"testing"
)

func TestSequentialSP(t *testing.T) {
	fileTestFloyd(t, SequentialSP)
}

func TestSequentialSPWithPath(t *testing.T) {
	fileTestFloydWithPath(t, SequentialSPWithPath)
}

func BenchmarkSequentialSP(b *testing.B) {
	fileBenchmarkFloyd(b, SequentialSP)
}
