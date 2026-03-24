package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

const INF = math.MaxInt

func SaveToFile(mat [][]int, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	n := len(mat)
	w := csv.NewWriter(file)
	row := make([]string, len(mat[0]))

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == INF {
				row[i] = "-"
			} else {
				row[i] = strconv.FormatInt(int64(mat[i][j]), 10)
			}
		}
		err := w.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}

func InputFromFile(filePath string) ([][]int, error) {
	var input [][]int

	file, err := os.Open(filePath)
	if err != nil {
		return input, err
	}
	defer file.Close()

	r := csv.NewReader(file)

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

	input = make([][]int, n)
	for i := 0; i < n; i++ {
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
				num, err := strconv.Atoi(row[j])
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

func ReadMatrixFromConsole() [][]int {
	var n int
	var mat [][]int

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
		if err == nil {
			break
		}
	}

	fmt.Println("Entering adjacency matrix rows separated by commas (\"-\" for absent edges)...")

	r := csv.NewReader(os.Stdin)
	mat = make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)

		for {
			fmt.Printf("Enter row #%d:\n", i+1)
			elems, err := r.Read()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if len(elems) != n {
				fmt.Println("Incorrect row length")
				continue
			}
			for j := 0; j < n; j++ {
				if elems[j] == "-" {
					mat[i][j] = INF
				} else {
					num, err := strconv.Atoi(elems[j])
					if err != nil {
						fmt.Println(err)
						break
					}
					mat[i][j] = num
				}
			}
			if err == nil {
				break
			}
		}
	}

	return mat
}
