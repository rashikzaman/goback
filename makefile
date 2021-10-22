install:
	cd cmd && go mod download
run:
	go run cmd/main.go
build:
	go build -o goback cmd/*.go