BINARY_DIR=bin/markdown2html
BINARY_NAME=markdown2html
check:
	@golangci-lint run ./...
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(BINARY_DIR)/$(BINARY_NAME)