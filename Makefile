.PHONY: ci deps test lint staticcheck

ci: deps test lint staticcheck

deps:
	go mod download

test:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

lint:
	golangci-lint run ./...

staticcheck:
	staticcheck -checks=all ./...
