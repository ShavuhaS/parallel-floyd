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

func InputFromFile(filePath string) ([][]float64, error) {
	var input [][]float64

	file, err := os.Open(filePath)
	if err != nil {
		return input, err
	}
	defer file.Close()

	r := csv.NewReader(file)

	return InputFromCsv(r)
}

func InputFromCsv(r *csv.Reader) ([][]float64, error) {
	var input [][]float64

	r.FieldsPerRecord = -1

	firstRow, err := r.Read()
	if err != nil {
		return input, err
	}
	if len(firstRow) != 1 {
		return input, errors.New("persistence: Invalid input format")
	}

	n, err := strconv.Atoi(firstRow[0])
	if err != nil {
		return input, err
	}

	input = make([][]float64, n)
	for i := 0; i < n; i++ {
		input[i] = make([]float64, n)

		row, err := r.Read()
		if err != nil {
			return input, err
		}
		if len(row) != n {
			return input, errors.New("persistence: Invalid input format")
		}
		for j := 0; j < n; j++ {
			if row[j] == "-" {
				input[i][j] = INF
			} else {
				num, err := strconv.ParseFloat(row[j], 64)
				if err != nil {
					return input, err
				}
				input[i][j] = num
			}
		}
	}

	_, err = r.Read()
	if err != io.EOF {
		return input, errors.New("persistence: Invalid input format (too much data)")
	}

	return input, nil
}

func SaveInputToFile(adjMat [][]float64, filePath string) error {
	n := len(adjMat)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	w.Write([]string{fmt.Sprint(n)})

	return SaveOutputToCsv(adjMat, w)
}

func SaveOutputToFile(mat [][]float64, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	return SaveOutputToCsv(mat, w)
}

func SaveOutputToCsv(mat [][]float64, w *csv.Writer) error {
	n := len(mat)
	row := make([]string, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == INF {
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

func OutputFromFile(filePath string) ([][]float64, error) {
	return InputFromFile(filePath)
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
