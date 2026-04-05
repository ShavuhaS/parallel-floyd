#!/bin/bash

go test -bench=^Bench.*BestConfig$ -run=^$ -benchtime 20x ./internal/floyd | tee benchresults_best.txt