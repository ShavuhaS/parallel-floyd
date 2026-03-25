package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ShavuhaS/parallel-floyd/internal/floyd"
	"github.com/ShavuhaS/parallel-floyd/internal/utils"
)

type Algorithm int

const (
	Sequential Algorithm = iota
	ParallelRowed
	ParallelBlocked
)

type FloydConfig struct {
	data         [][]float64
	algorithm    Algorithm
	inputFile    *string
	outputFile   *string
	routineCount int
	maxProcs     int
}

func (cfg *FloydConfig) Print() {
	fmt.Printf("%+v\n", *cfg)
	if cfg.inputFile != nil {
		fmt.Println("Input file:", *cfg.inputFile)
	}
	if cfg.outputFile != nil {
		fmt.Println("Output file:", *cfg.outputFile)
	}
}

func parseConfig() *FloydConfig {
	algo := flag.String("algorithm", "sequential", "Algorithm to use.")
	routineCount := flag.Int("routines", runtime.NumCPU(), "The amount of goroutines to use for parallel algorithms.")
	maxProcs := flag.Int("max-procs", runtime.NumCPU(), "The value of GOMAXPROCS env variable.")

	input := flag.String("input", "none", "File .txt from which to read the input (to which to write the input).")
	saveInput := flag.Bool("save-input", false, "Whether to save the input to a file.")
	output := flag.String("output", "none", "File (or dir) to which to save the output.")
	saveOutput := flag.Bool("save-output", false, "Whether to save the output to a file.")

	vertexCount := flag.Int("v", -1, "The number of vertices in a generated graph.")
	edgeProb := flag.Float64("edge-probability", 1.0, "Probability of an edge generation.")
	minEdge := flag.Float64("min-edge", 1.0, "Min edge length.")
	maxEdge := flag.Float64("max-edge", 15.0, "Max edge length.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", filepath.Base(os.Args[0]))

		fmt.Fprintf(os.Stderr, "Algorithm specific options:\n")
		fmt.Fprintf(os.Stderr, "\t-algorithm string\n\t\tAlgorithm to use.\n\t\tAvailable options: sequential, parallel-rowed, parallel-blocked. (default \"sequential\")\n")
		fmt.Fprintf(os.Stderr, "\t-routines int\n\t\tThe amount of goroutines to use. Defaults to the number of CPUs. (default %d)\n", runtime.NumCPU())
		fmt.Fprintf(os.Stderr, "\t-max-procs int\n\t\tThe value of GOMAXPROCS env variable. Defaults to the number of CPUs. (default %d)\n\n", runtime.NumCPU())

		fmt.Fprintf(os.Stderr, "I/O options:\n")
		fmt.Fprintf(os.Stderr, "\t-input string\n\t\tFile .txt or dir to which to write the input (with --save-input).\n\t\tFile .txt from which to read the input (without --save-input). The first line contains a single number \"n\" \n\t\tdenoting the number of vertices. Then it's followed by csv lines of the resulting matrix.\n\t\tAbsent edges should be denoted as \"-\". (default \"none\")\n")
		fmt.Fprintf(os.Stderr, "\t-save-input\n\t\tWhether to save the input to a file. (default false)\n")
		fmt.Fprintf(os.Stderr, "\t-output string\n\t\tFile (or dir) to which to save the output. (default \"none\")\n")
		fmt.Fprintf(os.Stderr, "\t-save-output\n\t\tWhether to save the output to a file. (default false)\n\n")

		fmt.Fprintf(os.Stderr, "Matrix generation options:\n")
		fmt.Fprintf(os.Stderr, "\t-v int\n\t\tThe number of vertices in a generated graph (if vertexCount is specified). (default -1)\n")
		fmt.Fprintf(os.Stderr, "\t-edge-probability float\n\t\tProbability of an edge generation. (default 1.0)\n")
		fmt.Fprintf(os.Stderr, "\t-min-edge float\n\t\tMin edge length. (default 1.0)\n")
		fmt.Fprintf(os.Stderr, "\t-max-edge float\n\t\tMax edge length. (default 15.0)\n")
	}

	flag.Parse()

	log.Default().SetFlags(0)

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

	if *input == "none" {
		input = nil
	} else {
		absPath, err := filepath.Abs(*input)
		if err != nil {
			log.Fatalln(err)
		}
		*input = absPath
	}

	if *output == "none" {
		output = nil
	} else {
		absPath, err := filepath.Abs(*output)
		if err != nil {
			log.Fatalln(err)
		}
		*output = absPath
	}

	if input != nil && *saveInput == false {
		mat, err := utils.InputFromFile(*input)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		cfg.data = mat
	} else if *vertexCount != -1 {
		if *vertexCount <= 0 {
			log.Fatalf("error: Invalid vertex count")
		}
		cfg.data = utils.GenerateMatrix(*vertexCount, *minEdge, *maxEdge, *edgeProb)

		if !(*saveInput) {
			fmt.Printf("Generated matrix (V = %v):\n", len(cfg.data))
			fmt.Println()
			fmt.Println(utils.GetMatrixString(cfg.data, floyd.INF))
			fmt.Println()
		}
	} else {
		cfg.data = utils.ReadInputFromConsole()
	}

	if *saveInput {
		if input == nil {
			p, err := getInputFilePath(len(cfg.data))
			if err != nil {
				log.Fatalln(err)
			}
			input = &p
		} else if isDir, err := utils.IsDir(*input); err != nil || isDir {
			if err != nil {
				log.Fatalln(err)
			}
			p := fmt.Sprintf("%v/%v", *input, generateInputFileName(len(cfg.data)))
			input = &p
		} else if filepath.Ext(*input) != ".txt" {
			log.Fatalln("Program only accepts txt files as input!")
		}
	}

	cfg.inputFile = input

	if *saveOutput || output != nil {
		var inputFilePath string
		if input == nil {
			inputFile, err := getInputFilePath(len(cfg.data))
			if err != nil {
				log.Fatalln(err)
			}
			inputFilePath = inputFile
		} else {
			inputFilePath = *input
		}

		if output == nil {
			outputFile := generateOutputFileName(inputFilePath, *algo)
			output = &outputFile
		} else if isDir, err := utils.IsDir(*output); err != nil || isDir {
			if err != nil {
				log.Fatalln(err)
			}
			inputPath := fmt.Sprintf("%v/%v", *output, filepath.Base(inputFilePath))
			outputFile := generateOutputFileName(inputPath, *algo)
			output = &outputFile
		}

	}

	cfg.outputFile = output

	cfg.routineCount = *routineCount
	cfg.maxProcs = *maxProcs

	return &cfg
}

func getInputFilePath(n int) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(wd, generateInputFileName(n))
	return path, nil
}

func generateInputFileName(n int) string {
	return fmt.Sprintf("input_%vv_%v.txt", n, time.Now().UnixMilli())
}

func generateOutputFileName(inputFileName string, algo string) string {
	ext := filepath.Ext(inputFileName)
	nameWithoutExt := inputFileName[:len(inputFileName)-len(ext)]
	return fmt.Sprintf("%v.%v.out%v", nameWithoutExt, algo, ext)
}
