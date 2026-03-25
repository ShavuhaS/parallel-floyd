package main

import (
	"fmt"
	"log"

	"github.com/ShavuhaS/parallel-floyd/internal/floyd"
	"github.com/ShavuhaS/parallel-floyd/internal/utils"
)

func main() {
	cfg := parseConfig()

	var dist [][]float64
	var prev [][]int
	switch cfg.algorithm {
	case Sequential:
		dist, prev = floyd.SequentialSPWithPath(cfg.data)
	default:
		log.Fatalln("Unknown algorithm type! Use --help to list available algorithms")
	}

	if cfg.outputDistFile == nil {
		fmt.Println("Results:")
		fmt.Println()
		for u := 0; u < len(cfg.data); u++ {
			for v := 0; v < len(cfg.data); v++ {
				path := floyd.GetShortestPath(prev, u, v)
				if len(path) == 0 {
					fmt.Printf("Path from %v to %v: no path (INF)\n", u, v)
				} else {
					fmt.Println(utils.GetPathString(cfg.data, path, dist[u][v]))
				}
			}
		}
		fmt.Println()
	}

	if cfg.inputFile != nil {
		err := utils.SaveInputToFile(cfg.data, *cfg.inputFile)
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
		fmt.Println("Input successfuly saved to", *cfg.inputFile)
	}

	if cfg.outputDistFile != nil {
		err := utils.SaveDistToFile(dist, *cfg.outputDistFile)
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
		fmt.Println("Output (dist) successfuly saved to", *cfg.outputDistFile)
	}

	if cfg.outputPrevFile != nil {
		err := utils.SavePrevToFile(prev, *cfg.outputPrevFile)
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
		fmt.Println("Output (prev) successfuly saved to", *cfg.outputPrevFile)
	}
}
