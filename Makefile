UNAME_S := $(shell uname -s | tr 'A-Z' 'a-z')
PROJ_NAME := convert-to-quicken-csv
BIN_DIR := bin/

build: build-$(UNAME_S)

all: build-windows build-linux build-darwin docs

clean:
	@rm -Rf bin

build-%:
	@echo Building for $(*)
	@echo GOOS=$(*) GOARCH=amd64 go build -o $(BIN_DIR)$(PROJ_NAME)-$(*)-amd64$(SUFFIX) $(PROJ_NAME).go && \
	GOOS=$(*) GOARCH=amd64 go build -o $(BIN_DIR)$(PROJ_NAME)-$(*)-amd64$(SUFFIX) $(PROJ_NAME).go;
	@if [ "$(*)" == "darwin" ]; then \
	echo GOOS=$(*) GOARCH=arm64 go build -o $(BIN_DIR)$(PROJ_NAME)-$(*)-arm64$(SUFFIX) $(PROJ_NAME).go && \
	GOOS=$(*) GOARCH=arm64 go build -o $(BIN_DIR)$(PROJ_NAME)-$(*)-arm64$(SUFFIX) $(PROJ_NAME).go && \
	echo lipo -create -output $(BIN_DIR)$(PROJ_NAME)-$(*)-universal$(SUFFIX) $(BIN_DIR)$(PROJ_NAME)-$(*)-arm64$(SUFFIX) $(BIN_DIR)$(PROJ_NAME)-$(*)-amd64$(SUFFIX) && \
	lipo -create -output $(BIN_DIR)$(PROJ_NAME)-$(*)-universal$(SUFFIX) $(BIN_DIR)$(PROJ_NAME)-$(*)-arm64$(SUFFIX) $(BIN_DIR)$(PROJ_NAME)-$(*)-amd64$(SUFFIX) && \
	rm -f $(BIN_DIR)$(PROJ_NAME)-$(*)-amd64$(SUFFIX) && \
	rm -f $(BIN_DIR)$(PROJ_NAME)-$(*)-arm64$(SUFFIX); \
	fi

docs:
ifeq (, $(shell which doctoc))
 $(error "No doctoc in $(PATH), consider installing to run 'docs' target")
endif
	@doctoc README.md