.DEFAULT_GOAL := build
HAS_UPX := $(shell command -v upx 2> /dev/null)
BIN_DIR := ./bin

# Version and ldflags
VERSION ?= $(shell git describe --tags --always --dirty=-dev 2>/dev/null || git rev-parse --short HEAD)
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: build
build:
	go build -ldflags="$(LDFLAGS)" -o ./feishu2md cmd/*.go
ifneq ($(and $(COMPRESS),$(HAS_UPX)),)
	upx -9 ./feishu2md
endif

.PHONY: test
test:
	go test ./...

.PHONY: server
server:
	go build -o ./feishu2md4web web/*.go

.PHONY: image
image:
	docker build -t feishu2md .

.PHONY: docker
docker:
	docker run -it --rm -p 8080:8080 feishu2md

.PHONY: clean
clean:  ## Clean build bundles
	rm -f ./feishu2md ./feishu2md4web
	rm -rf $(BIN_DIR)

.PHONY: format
format:
	gofmt -l -w .

.PHONY: all
all: build server
	@echo "Build all done"

# --- Cross-platform builds to bin/ ---

.PHONY: build-linux-amd64
build-linux-amd64:
	@mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -ldflags="$(LDFLAGS)" \
		-o $(BIN_DIR)/feishu2md-linux-amd64 ./cmd
ifneq ($(and $(COMPRESS),$(HAS_UPX)),)
	upx -9 $(BIN_DIR)/feishu2md-linux-amd64
endif

.PHONY: build-linux-arm64
build-linux-arm64:
	@mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 \
		go build -ldflags="$(LDFLAGS)" \
		-o $(BIN_DIR)/feishu2md-linux-arm64 ./cmd
ifneq ($(and $(COMPRESS),$(HAS_UPX)),)
	upx -9 $(BIN_DIR)/feishu2md-linux-arm64
endif

.PHONY: build-darwin-arm64
build-darwin-arm64:
	@mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 \
		go build -ldflags="$(LDFLAGS)" \
		-o $(BIN_DIR)/feishu2md-darwin-arm64 ./cmd

.PHONY: build-windows-amd64
build-windows-amd64:
	@mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
		go build -ldflags="$(LDFLAGS)" \
		-o $(BIN_DIR)/feishu2md-windows-amd64.exe ./cmd
ifneq ($(and $(COMPRESS),$(HAS_UPX)),)
	upx -9 $(BIN_DIR)/feishu2md-windows-amd64.exe || true
endif

.PHONY: build-bin
build-bin: build-linux-amd64 build-linux-arm64 build-darwin-arm64 build-windows-amd64

.PHONY: build-all
build-all: build build-bin

# --- Packaging ---
.PHONY: package-linux-amd64
package-linux-amd64: build-linux-amd64
	@mkdir -p $(BIN_DIR)
	tar -C $(BIN_DIR) -czf $(BIN_DIR)/feishu2md_$(VERSION)_linux-amd64.tar.gz feishu2md-linux-amd64

.PHONY: package-linux-arm64
package-linux-arm64: build-linux-arm64
	@mkdir -p $(BIN_DIR)
	tar -C $(BIN_DIR) -czf $(BIN_DIR)/feishu2md_$(VERSION)_linux-arm64.tar.gz feishu2md-linux-arm64

.PHONY: package-darwin-arm64
package-darwin-arm64: build-darwin-arm64
	@mkdir -p $(BIN_DIR)
	tar -C $(BIN_DIR) -czf $(BIN_DIR)/feishu2md_$(VERSION)_darwin-arm64.tar.gz feishu2md-darwin-arm64

.PHONY: package-windows-amd64
package-windows-amd64: build-windows-amd64
	@mkdir -p $(BIN_DIR)
	zip -j $(BIN_DIR)/feishu2md_$(VERSION)_windows-amd64.zip $(BIN_DIR)/feishu2md-windows-amd64.exe >/dev/null

.PHONY: package-all
package-all: package-linux-amd64 package-linux-arm64 package-darwin-arm64 package-windows-amd64
