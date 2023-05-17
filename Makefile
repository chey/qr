LDFLAGS ?= -s -w
BUILD_DIR ?= .build

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $$(PATH), visit: https://golangci-lint.run/usage/install/ for more information")
endif

.PHONY: docker lint build install lib test test-lib clean

lint:
	golangci-lint run ./...

build: $(BUILD_DIR)/bin
	go build -ldflags "$(LDFLAGS)" -o $</

docker:
	docker build -t qr --build-arg VERSION=docker .

install:
	go install -ldflags "$(LDFLAGS)"

test:
	go test ./...

lib: $(BUILD_DIR)/lib
	export C_INCLUDE_PATH=$<; go build -buildmode=c-shared -ldflags "$(LDFLAGS)" -o $</libqr.so ./internal...

test-lib: $(BUILD_DIR)/bin/main_test
	export LD_LIBRARY_PATH=$(BUILD_DIR)/lib; $<

$(BUILD_DIR)/bin/main_test: $(BUILD_DIR)/bin $(BUILD_DIR)/lib/libqr.so $(BUILD_DIR)/lib/libqr.h
	gcc test/main.c -o $</main_test -L"$(BUILD_DIR)/lib" -I"$(BUILD_DIR)/lib" -lqr

clean:
	rm -rf $(BUILD_DIR)

$(BUILD_DIR)/lib/libqr.so: lib

$(BUILD_DIR)/lib/libqr.h: lib

$(BUILD_DIR):
	mkdir -p .build

$(BUILD_DIR)/bin: $(BUILD_DIR)
	mkdir -p $@

$(BUILD_DIR)/lib: $(BUILD_DIR)
	mkdir -p $@
