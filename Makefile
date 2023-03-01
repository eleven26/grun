.PHONY: init
init:
	go build -modfile=tools/go.mod -o bin/gofumpt mvdan.cc/gofumpt
	go build -modfile=tools/go.mod -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: check
check:
	if [ ! -f bin/gofumpt ]; then make init; fi
	bin/golangci-lint run

FILES = $(shell find . -type f -name '*.go')

.PHONY: format
format:
	if [ ! -f bin/gofumpt ]; then make init; fi
	go mod tidy
	bin/gofumpt -w $(FILES)

.PHONY: test
test:
	go test ./... -cover

.PHONY: all
all:
	make check
	make format
	make test
