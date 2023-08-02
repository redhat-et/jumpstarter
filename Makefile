
jumpstarter: main.go pkg/drivers/jumpstarter-board/*.go pkg/harness/*.go cmd/*.go
	go build

fmt:
	gofmt -w -s .

all: jumpstarter

.PHONY: all fmt

