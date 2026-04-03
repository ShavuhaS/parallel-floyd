# Parallel Floyd-Warshall Algorithm in Go

ІП-31 Хохотва Андрій

## Build
Go version: 1.26

Using Make:

```bash
make build
```

Using Go compiler:

```bash
go build -o <dir> ./cmd/floyd
``` 

## Usage

### Algorithm specific options:
	-algorithm string
		Algorithm to use.
		Available options: sequential, parallel-rowed, parallel-blocked, parallel-phaseblocked. (default "sequential")
	-routines int
		The amount of goroutines to use. Defaults to the number of CPUs. (default 12)
	-max-procs int
		The value of GOMAXPROCS env variable. Defaults to the number of CPUs. (default 12)

### I/O options:
	-input string
		File .txt or dir to which to write the input (with --save-input).
		File .txt from which to read the input (without --save-input). The first line contains a single number "n" 
		denoting the number of vertices. Then it's followed by csv lines of the resulting matrix.
		Absent edges should be denoted as "-". (default "none")
	-save-input
		Whether to save the input to a file. (default false)
	-output string
		File name prefix (or dir) to which to save the output (dist and prev matrices). (default "none")
	-save-output
		Whether to save the output to a file. (default false)
	-print-output
		Whether to print the output to stdout. (default true)

### Matrix generation options:
	-v int
		The number of vertices in a generated graph (if vertexCount is specified). (default -1)
	-edge-probability float
		Probability of an edge generation. (default 1.0)
	-min-edge float
		Min edge length. (default 1.0)
	-max-edge float
		Max edge length. (default 15.0)
