GO ?= go
TARGET := bhyve-vm-goagent

all: build

build:
	@$(GO) build -o $(TARGET) $^

clean:
	@$(GO) clean
	@rm -rf $(TARGET)

.PHONY: all build clean
