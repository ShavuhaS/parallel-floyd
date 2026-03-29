#!/bin/bash
go test -v -timeout 20h ./internal/floyd/ | tee tests.txt

