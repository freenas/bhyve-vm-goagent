GO ?= go
TARGET := bhyve-vm-goagent
OS := freebsd netbsd linux windows
ARCH := 386 amd64

all: build

build:
	@for os in $(OS); do \
		for arch in $(ARCH); do \
		echo "===> building: $(TARGET)-$$os-$$arch"; \
		GOOS=$$os GOARCH=$$arch go build -o $(TARGET)-$$os-$$arch $^ ;\
		done \
	done \

clean:
	@$(GO) clean
	@for os in $(OS); do \
		for arch in $(ARCH); do \
		echo "===> Removing: $(TARGET)-$$os-$$arch"; \
		rm -f $(TARGET)-$$os-$$arch $^ ;\
		done \
	done \


.PHONY: all build clean
