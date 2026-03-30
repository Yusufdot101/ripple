.PHONY:build-api
build-api:
	@go build -o tmp/api ./cmd/api

.PHONY:build-grpc
build-grpc:
	@go build -o tmp/grpc ./cmd/grpc

.PHONY:run-api
run-api:
	@./tmp/api

.PHONY:run-grpc
run-grpc:
	@./tmp/grpc

.PHONY:run-dev
run-dev:
	@air

.PHONY:test
test:
	@go test -v ./...

.PHONY:test-coverage
test-coverage:
	@go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
