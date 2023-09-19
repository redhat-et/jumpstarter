
jumpstarter: main.go pkg/drivers/jumpstarter-board/*.go pkg/runner/* pkg/harness/*.go cmd/*.go pkg/tools/*.go
	go build

fmt:
	gofmt -w -s .

all: jumpstarter

.PHONY: all fmt

