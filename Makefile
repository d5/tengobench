vet:
	@go vet ./...

start: vet
	@go run ./cmd/bench ./code

fmt:
	@go fmt ./...
