package floyd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ShavuhaS/parallel-floyd/internal/utils"
)

const TEST_INPUT_DIR = "testdata/input"
const TEST_DIST_DIR = "testdata/dist"

// const TEST_PREV_DIR = "testdata/prev"

func fileTestFloyd(t *testing.T, floyd func([][]float64) [][]float64) {
	testInputs, err := os.ReadDir(TEST_INPUT_DIR)
	if err != nil {
		t.Fatalf("Unable to read test inputs directory: %v", err)
	}

	for _, file := range testInputs {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".input.txt") {
			continue
		}

		baseName := strings.TrimSuffix(file.Name(), ".input.txt")

		t.Run(baseName, func(t *testing.T) {
			inputPath := filepath.Join(TEST_INPUT_DIR, baseName+".input.txt")
			distPath := filepath.Join(TEST_DIST_DIR, baseName+".dist.txt")

			inputMat, err := utils.InputFromFile(inputPath)
			if err != nil {
				t.Fatalf("Unable to read test input file: %v", err)
			}

			expectedDist, err := utils.DistFromFile(distPath)
			if err != nil {
				t.Fatalf("Unable to read test output file (dist): %v", err)
			}

			actualDist := floyd(inputMat)

			utils.AssertMatricesEqual(t, actualDist, expectedDist, INF)
		})
	}
}

func fileTestFloydWithPath(t *testing.T, floyd func([][]float64) ([][]float64, [][]int)) {
	testInputs, err := os.ReadDir(TEST_INPUT_DIR)
	if err != nil {
		t.Fatalf("Unable to read test inputs directory: %v", err)
	}

	for _, file := range testInputs {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".input.txt") {
			continue
		}

		baseName := strings.TrimSuffix(file.Name(), ".input.txt")

		t.Run(baseName, func(t *testing.T) {
			inputPath := filepath.Join(TEST_INPUT_DIR, baseName+".input.txt")
			distPath := filepath.Join(TEST_DIST_DIR, baseName+".dist.txt")
			// prevPath := filepath.Join(TEST_PREV_DIR, baseName+".prev.txt")

			inputMat, err := utils.InputFromFile(inputPath)
			if err != nil {
				t.Fatalf("Unable to read test input file: %v", err)
			}

			expectedDist, err := utils.DistFromFile(distPath)
			if err != nil {
				t.Fatalf("Unable to read test output file (dist): %v", err)
			}

			// expectedPrev, err := utils.PrevFromFile(prevPath)
			// if err != nil {
			// 	t.Fatalf("Unable to read test output file (prev): %v", err)
			// }

			actualDist, actualPrev := floyd(inputMat)

			utils.AssertMatricesEqual(t, actualDist, expectedDist, INF)
			// utils.AssertMatricesEqual(t, actualPrev, expectedPrev, math.MaxInt)
			utils.AssertFloydDistMatchesPrev(t, inputMat, actualDist, actualPrev)
		})
	}
}
