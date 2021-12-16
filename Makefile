all: build run

build:
	go build -o bin/mutorere cmd/mutorere/main.go

run:
	go run cmd/mutorere/main.go