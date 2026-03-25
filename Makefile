.PHONY: build

build: $(wildcard cmd/*.go) $(wildcard internal/*.go)
	go build -o ./bin/floyd ./cmd/floyd/