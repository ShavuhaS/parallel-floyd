package floyd

import (
	"cmp"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"

	"github.com/ShavuhaS/parallel-floyd/internal/utils"
)

const TEST_INPUT_DIR = "testdata/input"
const TEST_DIST_DIR = "testdata/dist"

func fileTestFloyd(t *testing.T, floyd func([][]float64)) {
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

			dist := InitDist(inputMat)
			floyd(dist)

			utils.AssertMatricesEqual(t, dist, expectedDist, INF)
		})
	}
}

func fileTestFloydWithPath(t *testing.T, floyd func([][]float64, [][]int)) {
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

			dist := InitDist(inputMat)
			prev := InitPrev(inputMat)
			floyd(dist, prev)

			utils.AssertMatricesEqual(t, dist, expectedDist, INF)
			utils.AssertFloydDistMatchesPrev(t, inputMat, dist, prev)
		})
	}
}

func fileBenchmarkFloyd(b *testing.B, floyd func([][]float64)) {
	testInputs, err := os.ReadDir(TEST_INPUT_DIR)
	if err != nil {
		b.Fatalf("Unable to read test inputs directory: %v", err)
	}

	slices.SortFunc(testInputs, func(a, b os.DirEntry) int {
		var vA, vB int
		fmt.Sscanf(a.Name(), "testcase_%dv", &vA)
		fmt.Sscanf(b.Name(), "testcase_%dv", &vB)
		return cmp.Compare(vA, vB)
	})

	for _, file := range testInputs {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".input.txt") {
			continue
		}

		var vertexCount int
		_, err := fmt.Sscanf(file.Name(), "testcase_%dv.input.txt", &vertexCount)
		if err != nil {
			b.Fatal(err)
		}

		benchName := fmt.Sprintf("V=%d", vertexCount)

		b.Run(benchName, func(b *testing.B) {
			inputPath := filepath.Join(TEST_INPUT_DIR, file.Name())

			inputMat, err := utils.InputFromFile(inputPath)
			if err != nil {
				b.Fatalf("Unable to read test input file: %v", err)
			}

			b.ResetTimer()
			for range b.N {
				b.StopTimer()

				runtime.GC()
				dist := InitDist(inputMat)

				b.StartTimer()
				floyd(dist)
			}
		})
	}
}

func benchmarkParallelFloydGoroutines(b *testing.B, parallelFloyd func([][]float64, int), v int) {
	goroutineCounts := []int{2, 4, 9, 12, 24, 48, 144, 1000, 5000}

	fileName := fmt.Sprintf("testcase_%dv.input.txt", v)
	inputPath := filepath.Join(TEST_INPUT_DIR, fileName)

	inputMat, err := utils.InputFromFile(inputPath)
	if err != nil {
		b.Fatalf("Unable to read test input file: %v", err)
	}

	originalMaxProcs := runtime.GOMAXPROCS(0)
	defer runtime.GOMAXPROCS(originalMaxProcs)

	numCpu := runtime.GOMAXPROCS(runtime.NumCPU())

	for i := range goroutineCounts {
		routines := goroutineCounts[i]
		benchName := fmt.Sprintf("V=%d_GR=%d_MaxProcs=%d", v, routines, numCpu)

		b.Run(benchName, func(b *testing.B) {
			b.ResetTimer()
			for range b.N {
				b.StopTimer()

				runtime.GC()
				dist := InitDist(inputMat)

				b.StartTimer()
				parallelFloyd(dist, routines)
			}
		})
	}
}

func benchmarkParallelFloydProcs(b *testing.B, parallelFloyd func([][]float64, int), v int) {
	const routineCount = 48
	goMaxProcs := []int{1, 2, 4, 8, 12, 16, 24, 32}

	fileName := fmt.Sprintf("testcase_%dv.input.txt", v)
	inputPath := filepath.Join(TEST_INPUT_DIR, fileName)

	inputMat, err := utils.InputFromFile(inputPath)
	if err != nil {
		b.Fatalf("Unable to read test input file: %v", err)
	}

	for i := range goMaxProcs {
		maxProcs := goMaxProcs[i]
		benchName := fmt.Sprintf("V=%d_GR=%d_MaxProcs=%d", v, routineCount, maxProcs)

		b.Run(benchName, func(b *testing.B) {
			originalMaxProcs := runtime.GOMAXPROCS(0)
			defer runtime.GOMAXPROCS(originalMaxProcs)

			runtime.GOMAXPROCS(maxProcs)
			b.ResetTimer()
			for range b.N {
				b.StopTimer()

				runtime.GC()
				dist := InitDist(inputMat)

				b.StartTimer()
				parallelFloyd(dist, routineCount)
			}
		})
	}
}
