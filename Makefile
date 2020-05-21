unit-test:
	go test ./... -v -race -count 100 -tags=unit
integration-test:
	go test ./... -v -race -tags=integration
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
	docker-compose -f deployments/docker-compose.yml up
run-dev:
	docker-compose -f deployments/docker-compose.yml -f deployments/docker-compose.dev.yml up --build
run-integration-test:
	docker-compose -f deployments/docker-compose.yml -f deployments/docker-compose.test.yml up --build