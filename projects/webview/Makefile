OS = $(shell uname | tr A-Z a-z)

BUILD_DIR ?= build
VERSION ?= $(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
DATE_FMT = +%FT%T%z
ifdef SOURCE_DATE_EPOCH
    BUILD_DATE ?= $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u -r "$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u "$(DATE_FMT)")
else
    BUILD_DATE ?= $(shell date "$(DATE_FMT)")
endif
LDFLAGS += -X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildDate=${BUILD_DATE}
export CGO_ENABLED ?= 1
ifeq (${VERBOSE}, 1)
ifeq ($(filter -v,${GOARGS}),)
	GOARGS += -v
endif
TEST_FORMAT = short-verbose
endif

-include override.mk

.PHONY: clean
clean: ## Clean builds
	rm -rf ${BUILD_DIR}/
	rm -rf out

.PHONY: goversion
goversion:
ifneq (${IGNORE_GOLANG_VERSION_REQ}, 1)
	@printf "${GOLANG_VERSION}\n$$(go version | awk '{sub(/^go/, "", $$3);print $$3}')" | sort -t '.' -k 1,1 -k 2,2 -k 3,3 -g | head -1 | grep -q -E "^${GOLANG_VERSION}$$" || (printf "Required Go version is ${GOLANG_VERSION}\nInstalled: `go version`" && exit 1)
endif

.PHONY: build
build: goversion
ifeq (${VERBOSE}, 1)
	go env
endif

	@mkdir -p ${BUILD_DIR}
	go build ${GOARGS} -tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/ ./npn-webview.go

.PHONY: build-release
build-release: goversion ## Build all binaries without debug information
	@env GOOS=${GOOS} GOARCH=${GOARCH} ${MAKE} LDFLAGS="-w ${LDFLAGS}" GOARGS="${GOARGS} -trimpath" BUILD_DIR="${BUILD_DIR}/release" build

.PHONY: build-debug
build-debug: goversion compile-templates ## Build all binaries with remote debugging capabilities
	@${MAKE} GOARGS="${GOARGS} -gcflags \"all=-N -l\"" BUILD_DIR="${BUILD_DIR}/debug" build

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
