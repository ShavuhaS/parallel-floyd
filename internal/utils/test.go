package utils

import (
	"testing"
)

func AssertMatricesEqual[T comparable](t *testing.T, got [][]T, expected [][]T, infValue T) {
	if len(got) != len(expected) {
		failMatricesEqual(t, got, expected, infValue)
		return
	}

	for i := 0; i < len(got); i++ {
		if len(got[i]) != len(expected[i]) {
			failMatricesEqual(t, got, expected, infValue)
			return
		}
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != expected[i][j] {
				failMatricesEqual(t, got, expected, infValue)
				return
			}
		}
	}
}

func failMatricesEqual[T comparable](t *testing.T, got [][]T, expected [][]T, infValue T) {
	gotString := GetMatrixString(got, infValue)
	expectedString := GetMatrixString(expected, infValue)
	t.Errorf("Matrices are not equal!\nGot:\n%v\n\nExpected:\n%v\n", gotString, expectedString)
}

func AssertFloydDistMatchesPrev(t *testing.T, adjMat, dist [][]float64, prev [][]int) {
	n := len(adjMat)

	for i := range n {
		for j := range n {
			path := GetShortestPath(prev, i, j)
			pathLength := float64(0)

			for v := 1; v < len(path); v++ {
				pathLength += adjMat[path[v-1]][path[v]]
			}

			if pathLength != dist[i][j] {
				t.Errorf("Invalid path from %v to %v:\nExpected SP length: %v\nGot path with length: %v", i, j, dist[i][j], pathLength)
			}
		}
	}
}
