all: .format

.PHONY: all

.build: .format
	go build .

.format:
	go mod tidy
	go generate
	gofmt -w .
	goimports -w .
	golint ./...
	go-xray -d .