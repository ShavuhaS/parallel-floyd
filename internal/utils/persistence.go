package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const INF = math.MaxFloat64

func MatrixFromCsv[T comparable](r *csv.Reader, parse func(string) (T, error)) ([][]T, error) {
	r.FieldsPerRecord = -1

	firstRow, err := r.Read()
	if err != nil {
		return [][]T{}, err
	}
	if len(firstRow) != 1 {
		return [][]T{}, errors.New("persistence: Invalid input format")
	}

	n, err := strconv.Atoi(firstRow[0])
	if err != nil {
		return [][]T{}, err
	}

	input := make([][]T, n)
	for i := 0; i < n; i++ {
		input[i] = make([]T, n)

		row, err := r.Read()
		if err != nil {
			return input, err
		}
		if len(row) != n {
			return input, errors.New("persistence: Invalid input format")
		}
		for j := 0; j < n; j++ {
			val, err := parse(row[j])
			if err != nil {
				return [][]T{}, err
			}
			input[i][j] = val
		}
	}

	_, err = r.Read()
	if err != io.EOF {
		return input, errors.New("persistence: Invalid input format (too much data)")
	}

	return input, nil
}

func Float64MatrixFromCsv(r *csv.Reader) ([][]float64, error) {
	return MatrixFromCsv(r, func(s string) (float64, error) {
		if s == "-" {
			return INF, nil
		}
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		return num, nil
	})
}

func IntMatrixFromCsv(r *csv.Reader) ([][]int, error) {
	return MatrixFromCsv(r, func(s string) (int, error) {
		if s == "-" {
			return math.MaxInt, nil
		}
		num, err := strconv.ParseInt(s, 0, 32)
		if err != nil {
			return 0, err
		}
		return int(num), err
	})
}

func MatrixFromFile[T comparable](filePath string, parseCsv func(*csv.Reader) ([][]T, error)) ([][]T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return [][]T{}, err
	}
	defer file.Close()

	r := csv.NewReader(file)

	return parseCsv(r)
}

func MatrixToCsv[T comparable](w *csv.Writer, mat [][]T, noValue T) error {
	n := len(mat)
	row := make([]string, n)

	w.Write([]string{fmt.Sprint(n)})

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == noValue {
				row[j] = "-"
			} else {
				row[j] = fmt.Sprintf("%v", mat[i][j])
			}
		}
		err := w.Write(row)
		if err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}

func MatrixToFile[T comparable](filePath string, mat [][]T, writeCsv func(*csv.Writer, [][]T) error) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	return writeCsv(w, mat)
}

func Float64MatrixToCsv(w *csv.Writer, mat [][]float64) error {
	return MatrixToCsv(w, mat, INF)
}

func IntMatrixToCsv(w *csv.Writer, mat [][]int) error {
	return MatrixToCsv(w, mat, math.MaxInt)
}

func SaveInputToFile(adjMat [][]float64, filePath string) error {
	return MatrixToFile(filePath, adjMat, Float64MatrixToCsv)
}

func SaveDistToFile(distMat [][]float64, filePath string) error {
	return MatrixToFile(filePath, distMat, Float64MatrixToCsv)
}

func SavePrevToFile(prevMat [][]int, filePath string) error {
	n := len(prevMat)

	prev := make([][]int, n)
	for i := 0; i < n; i++ {
		prev[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if prev[i][j] == -1 {
				prev[i][j] = math.MaxInt
			} else {
				prev[i][j] = prevMat[i][j]
			}
		}
	}

	return MatrixToFile(filePath, prev, IntMatrixToCsv)
}

func InputFromFile(filePath string) ([][]float64, error) {
	return MatrixFromFile(filePath, Float64MatrixFromCsv)
}

func DistFromFile(filePath string) ([][]float64, error) {
	return MatrixFromFile(filePath, Float64MatrixFromCsv)
}

func PrevFromFile(filePath string) ([][]int, error) {
	prev, err := MatrixFromFile(filePath, IntMatrixFromCsv)
	if err != nil {
		return [][]int{}, nil
	}

	for i := 0; i < len(prev); i++ {
		for j := 0; j < len(prev); j++ {
			if prev[i][j] == math.MaxInt {
				prev[i][j] = -1
			}
		}
	}

	return prev, nil
}

func ReadInputFromConsole() [][]float64 {
	var n int
	var mat [][]float64

	for {
		fmt.Printf("Enter vertex count: ")
		_, err := fmt.Scanf("%d\n", &n)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n <= 0 {
			fmt.Println("error: Vertex count must be positive")
			continue
		}
		break
	}

	fmt.Println("Entering adjacency matrix rows separated by commas (\"-\" for absent edges)...")

	r := csv.NewReader(os.Stdin)
	r.FieldsPerRecord = -1

	mat = make([][]float64, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]float64, n)

		valid := false
		for !valid {
			valid = inputRowFromConsole(r, mat, i)
		}
	}

	return mat
}

func inputRowFromConsole(r *csv.Reader, mat [][]float64, i int) bool {
	fmt.Printf("Enter row #%d:\n", i+1)

	elems, err := r.Read()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if len(elems) != len(mat) {
		fmt.Println("Incorrect row length")
		return false
	}

	for j := 0; j < len(mat); j++ {
		elem := strings.TrimSpace(elems[j])
		if i == j && elem != "0" {
			fmt.Println("Main diagonal must be nulls!")
			return false
		}

		if elem == "-" {
			mat[i][j] = INF
		} else {
			num, err := strconv.ParseFloat(elem, 64)
			if err != nil {
				fmt.Println(err)
				return false
			}
			mat[i][j] = num
		}
	}

	return true
}
