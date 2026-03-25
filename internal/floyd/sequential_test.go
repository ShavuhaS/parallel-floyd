package floyd

import (
	"testing"
)

func TestSequentialSP(t *testing.T) {
	fileTestFloyd(t, SequentialSP)
}

func TestFloydWarshallWithPath(t *testing.T) {
	fileTestFloydWithPath(t, SequentialSPWithPath)
}
