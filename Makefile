test:
	go test ./... -v -race -count 100
run-go:
	go run cmd/main.go
.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/app cmd/main.go
build-dev:
	go build -o bin/app cmd/main.go
generate:
	go generate cmd/main.go
install:
	go install
run-prod:
	docker-compose -f deployments/docker-compose.yml up --build
run-dev:
	docker-compose -f deployments/docker-compose.yml -f deployments/docker-compose.dev.yml up --build