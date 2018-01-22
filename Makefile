GO ?= go
TARGET := bhyve-vm-goagent
OS := freebsd netbsd linux windows
ARCH := 386 amd64

all: build

deps:
	@echo "===> Downloading crossbuild dependencies."
	go get github.com/gorilla/websocket
	go get github.com/shirou/w32
	go get github.com/StackExchange/wmi
	go get github.com/go-ole/go-ole
	go get github.com/go-ole/go-ole/oleutil

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


.PHONY: all deps build clean
