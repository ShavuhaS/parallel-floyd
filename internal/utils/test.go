package utils

import (
	"testing"
)

func AssertMatricesEqual(t *testing.T, got [][]int, expected [][]int) {
	if len(got) != len(expected) {
		failMatricesEqual(t, got, expected)
		return
	}

	for i := 0; i < len(got); i++ {
		if len(got[i]) != len(expected[i]) {
			failMatricesEqual(t, got, expected)
			return
		}
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != expected[i][j] {
				failMatricesEqual(t, got, expected)
				return
			}
		}
	}
}

func failMatricesEqual(t *testing.T, got [][]int, expected [][]int) {
	gotString := GetMatrixString(got, INF)
	expectedString := GetMatrixString(expected, INF)
	t.Errorf("Matrices are not equal!\nGot:\n%v\n\nExpected:\n%v\n", gotString, expectedString)
}
