
jumpstarter: main.go pkg/drivers/jumpstarter-board/*.go pkg/harness/*.go cmd/*.go
	go build

all: jumpstarter

