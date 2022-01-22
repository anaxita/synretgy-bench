lint:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.43.0 golangci-lint run

bench:
	go test -v ./... -bench . -count 1