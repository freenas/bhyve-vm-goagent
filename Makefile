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
                    if [ $$os == "openbsd" ] && [ $$arch == "amd64" ] ; then \
                        CGO_ENABLED="1" ; \
                        CC="gcc7" ; \
                        CXX="g++7" ; \
                    else if [ $$os == "openbsd" ] && [ $$arch == "386" ] ; then \
                        continue ; \
                    else \
                        CGO_ENABLED="0" ; \
                        CC="clang" ; \
                        CXX="clang++" ; \
                    fi ; \
                    fi ; \
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
