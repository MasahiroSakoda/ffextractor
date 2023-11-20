SHELL := /bin/bash

BIN_NAME := ffextractor
VERSION  := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)

GOOS      := $(shell go env GOOS)
GOARCH    := $(shell go env GOARCH)
PLATFORMS := darwin linux

GOBIN     := go
GOBUILD   := ${GOBIN} build
GOCLEAN   := ${GOBIN} clean
GOTEST    := ${GOBIN} test
GOGET     := ${GOBIN} get
GOINSTALL := ${GOBIN} install
GOTOOL    := ${GOBIN} tool

BUILD_DIR    := build
COVER_CMD    := ${GOTOOL} cover
COVER_PROFLE := coverage.out
COVER_TMP    := coverage.out.tmp
COVER_HTML   := coverage.html

GO_FILES  := $(shell find . -type f -name '*.go' -print)
PREFIX    := /usr/local
DIST_PATH := ${PREFIX}/${BIN_NAME}

.PHONY: all
all: deps install clean

.PHONY: deps
deps:
	${GOGET} -d -v ./...

.PHONY: build
build:
	${GOBIN} build .

.PHONY: build-all
# TODO: build for all OS & arch
build-all:

.PHONY: install
install: build
	cp ${BIN_NAME} ${DIST_PATH}

.PHONY: run
run: build
	./${BIN_NAME}

.PHONY: test
test:
	${GOTEST} -v ./...

.PHONY: coverage
coverage:
	rm -rf *.out ${COVER_HTML}
	${GOTEST} -cover -covermode=atomic ./... -coverprofile=${COVER_TMP}
	cat ${COVER_TMP} | grep -v "**_mock.go" | grep -v "wire_gen.go" > ${COVER_PROFLE}
	rm ${COVER_TMP}
	${COVER_CMD} -html=${COVER_PROFLE} -o ${COVER_HTML}
	open ${COVER_HTML}

.PHONY: clean
clean:
	${GOCLEAN} -v ./...
	rm -rf ${BUILD_DIR}

.PHONY: tag
tag:
	git tag "v$(VERSION)" --force

.PHONY: compress
# TODO: compress binary
compress:

.PHONY: release
# TODO: release with `goreleaser`
release: deps build-all compress
	goreleaser release --snapshot --clean
