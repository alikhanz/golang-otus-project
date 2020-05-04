test:
	go test ./... -v -race -count 100
run:
	go run cmd/main.go
.PHONY: build
build:
	go build -o bin/app cmd/main.go
generate:
	go generate cmd/main.go
install:
	go install