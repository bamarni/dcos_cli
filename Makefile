CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
PKG_DIR=/go/src/github.com/dcos/dcos-cli
BINARY_NAME=dcos

windows_EXE=.exe

.PHONY: default
default:
	@make $(shell uname | tr [A-Z] [a-z])

.PHONY: darwin linux windows
darwin linux windows:
	docker run \
        -v $(CURRENT_DIR):$(PKG_DIR) \
        -w $(PKG_DIR) \
        --privileged \
        --rm \
        golang:1.9.2 \
        bash -c "env GOOS=$(@) go build -o build/$(@)/$(BINARY_NAME)$($(@)_EXE) .../cmd/dcos"

.PHONY: vendor
vendor:
	docker run \
      -v $(CURRENT_DIR):$(PKG_DIR) \
      -w $(PKG_DIR) \
      --privileged \
      --rm \
      golang:1.9.2 \
      bash -c "go get -u github.com/golang/dep/... && dep ensure"

.PHONY: clean
clean:
	rm -rf build
