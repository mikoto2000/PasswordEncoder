APP_NAME := password-encoder
BUILD_DIR := ./build
MAIN_PKG := .

OSES := linux darwin windows
ARCHES := amd64 arm64

.PHONY: all build clean

all: build

build:
	@mkdir -p $(BUILD_DIR)
	@set -e; \
	for os in $(OSES); do \
		for arch in $(ARCHES); do \
			output="$(BUILD_DIR)/$(APP_NAME)-$${os}-$${arch}"; \
			if [ "$${os}" = "windows" ]; then \
				output="$${output}.exe"; \
			fi; \
			echo "Building $${output}"; \
			GOOS=$${os} GOARCH=$${arch} go build -o "$${output}" $(MAIN_PKG); \
		done; \
	done

clean:
	rm -rf $(BUILD_DIR)
