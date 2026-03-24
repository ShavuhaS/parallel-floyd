package main

import (
	"fmt"
	"log"

	"github.com/ShavuhaS/parallel-floyd/internal/floyd"
	"github.com/ShavuhaS/parallel-floyd/internal/utils"
)

func main() {
	cfg := parseConfig()

	var res [][]int
	switch cfg.algorithm {
	case Sequential:
		res = floyd.SequentialSP(cfg.data)
	default:
		log.Fatalln("Unknown algorithm type! Use --help to list available algorithms")
	}

	if cfg.saveOutput {
		err := utils.SaveToFile(res, cfg.outputFile)
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
	} else {
		fmt.Println("Result:")
		fmt.Println(utils.GetMatrixString(res, floyd.INF))
	}
}
