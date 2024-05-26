# Go files to format
BIN = terraform-teeworlds
GOFMT_FILES ?= $(shell find . -name "*.go")

default: fmt

fmt:
	gofmt -w $(GOFMT_FILES)

build:
	go build -v -o $(BIN)

clean:
	go clean -testcache

fclean: clean
	$(RM) $(BIN)

.PHONY: \
	fmt \
	clean \
	fclean
