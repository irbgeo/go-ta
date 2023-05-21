lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint cache clean
	golangci-lint run --fix -v

tests:
	go test ./...