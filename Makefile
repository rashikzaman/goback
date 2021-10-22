install:
	go mod download

run:
	go run main.go

build:
	go build

test:
	go test ./...