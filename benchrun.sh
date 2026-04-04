#!/bin/bash

go test -bench=. -run=^$ -benchtime 20x ./internal/floyd | tee benchresults.txt