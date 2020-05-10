.PHONY: build test
build: 
	go build -v ./cmd/apiserver
test:
	go test -v -race -timeout 10s ./...

.DEFAULT_GOAL :=build 