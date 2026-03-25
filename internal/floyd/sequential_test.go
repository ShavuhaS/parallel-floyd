package floyd

import (
	"math"
	"testing"

	"github.com/ShavuhaS/parallel-floyd/internal/utils"
)

func TestSequentialSP(t *testing.T) {
	for _, testCase := range floydTestCases {
		name, input, expected := testCase.name, testCase.input, testCase.expectedDist
		t.Run(name, func(t *testing.T) {
			got := SequentialSP(input)
			utils.AssertMatricesEqual(t, got, expected, INF)
		})
	}
}

func TestSequentialSPWithPath(t *testing.T) {
	for _, testCase := range floydTestCases {
		name := testCase.name
		input := testCase.input
		expectedPrev := testCase.expectedPrev
		expectedDist := testCase.expectedDist
		t.Run(name, func(t *testing.T) {
			gotDist, gotPrev := SequentialSPWithPath(input)
			utils.AssertMatricesEqual(t, gotDist, expectedDist, INF)
			utils.AssertMatricesEqual(t, gotPrev, expectedPrev, math.MaxInt)
		})
	}
}
