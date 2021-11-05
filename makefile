install:
	go mod download

run:
	go run main.go

run-h:
	gin --port 8080 run main.go

build:
	go build

test:
	go test ./...
