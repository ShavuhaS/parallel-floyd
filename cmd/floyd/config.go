package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/ShavuhaS/parallel-floyd/internal/utils"
)

type Algorithm int

const (
	Sequential Algorithm = iota
	ParallelRowed
	ParallelBlocked
)

type FloydConfig struct {
	data         [][]int
	algorithm    Algorithm
	outputFile   string
	saveOutput   bool
	routineCount int
	maxProcs     int
}

func parseConfig() *FloydConfig {
	algo := flag.String("algorithm", "sequential", `
Algorithm to use.
Available options: sequential, parallel-rowed, parallel-blocked.`)
	inputFile := flag.String("inputFile", "none", `
File from which to read the input. The first line contains a single number "n" denoting the number of vertices. Then it's followed by csv lines of the resulting matrix. Absent edges should be denoted as "-"`)
	outputFile := flag.String("outputFile", "none", "File to which to save the output.")
	vertexCount := flag.Int("v", -1, "The number of vertices in a generated graph (if vertexCount is specified).")
	edgeProb := flag.Float64("edgeProbability", 1.0, "Probability of an edge generation.")
	routineCount := flag.Int("routines", runtime.NumCPU(), `
The amount of goroutines to use. Defaults to the number of CPUs on the machine.`)
	maxProcs := flag.Int("maxProcs", runtime.NumCPU(), `
The value of GOMAXPROCS env variable. Defaults to the number of CPUs on the machine.`)

	flag.Parse()

	var cfg FloydConfig

	switch *algo {
	case "sequential":
		cfg.algorithm = Sequential
	case "parallel-rowed":
		cfg.algorithm = ParallelRowed
	case "parallel-blocked":
		cfg.algorithm = ParallelBlocked
	default:
		log.Fatalln("error: Unknown algorithm! See --help for available options")
	}

	if *inputFile != "none" {
		mat, err := utils.InputFromFile(*inputFile)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		cfg.data = mat
	} else if *vertexCount != -1 {
		if *vertexCount <= 0 {
			log.Fatalf("error: Invalid vertex count")
		}
		cfg.data = utils.GenerateMatrix(*vertexCount, 1, 15, *edgeProb)
	} else {
		cfg.data = utils.ReadMatrixFromConsole()
	}

	if *outputFile == "none" {

	}

	return &cfg
}
