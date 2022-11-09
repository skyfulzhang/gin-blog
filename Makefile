.PHONY:default install vat test build run  tool  clean  help

GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
BINARY_NAME=gin-blog
BINARY_UNIX=$(BINARY_NAME)_unix


all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

check:
	@go fmt ./
	@go vet ./

test:
	$(GOTEST) -v ./...

cover:
	@go test -coverprofile cover.out
	@go tool cover -html=cover.out

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make run: go  run ./main.go"
	@echo "make clean: remove object files and cached files"