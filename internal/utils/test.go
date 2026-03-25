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
