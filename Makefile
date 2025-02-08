DIR := $(PWD)
GOCMD=go

ARTIFACT_NAME=lagoon
ARTIFACT_DESTINATION ?= $(GOPATH)/bin

PKG=github.com/uselagoon/lagoon-cli
PKGMODPATH=$(DIR)/vendor

VERSION=$(shell echo $(shell git describe --abbrev=0 --tags)+$(shell git rev-parse --short=8 HEAD))
BUILD=$(shell date +%FT%T%z)

DOCKER_GO_VER=1.23
GO_VER=$(shell go version)
LDFLAGS=-w -s -X ${PKG}/cmd.lagoonCLIVersion=${VERSION} -X ${PKG}/cmd.lagoonCLIBuild=${BUILD}

GIT_ORIGIN=origin

all: deps test build docs
all-docker-linux: deps-docker test-docker build-docker-linux
all-docker-darwin: deps-docker test-docker build-docker-darwin

gen: deps
	GO111MODULE=on $(GOCMD) generate ./...
deps: 
	GO111MODULE=on ${GOCMD} get -v
test: gen
	GO111MODULE=on $(GOCMD) fmt ./...
	GO111MODULE=on $(GOCMD) vet ./...
	GO111MODULE=on $(GOCMD) test -v ./...

clean:
	$(GOCMD) clean

build: test
	GO111MODULE=on CGO_ENABLED=0 $(GOCMD) build -ldflags '${LDFLAGS} -X "${PKG}/cmd.lagoonCLIBuildGoVersion=${GO_VER}"' -o ${ARTIFACT_DESTINATION}/${ARTIFACT_NAME} -v
build-linux: test
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -ldflags '${LDFLAGS} -X "${PKG}/cmd.lagoonCLIBuildGoVersion=${GO_VER}"' -o builds/lagoon-cli-${VERSION}-linux-amd64 -v
build-darwin: test
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOCMD) build -ldflags '${LDFLAGS} -X "${PKG}/cmd.lagoonCLIBuildGoVersion=${GO_VER}"' -o builds/lagoon-cli-${VERSION}-darwin-amd64 -v
build-darwin-arm64: test
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOCMD) build -ldflags '${LDFLAGS} -X "${PKG}/cmd.lagoonCLIBuildGoVersion=${GO_VER}"' -o builds/lagoon-cli-${VERSION}-darwin-arm64 -v

docs: test
	LAGOON_GEN_DOCS=true GO111MODULE=on $(GOCMD) run main.go --docs 

tidy:
	GO111MODULE=on $(GOCMD) mod tidy

## build using docker golang
deps-docker:
	docker run \
	-v $(PKGMODPATH):/go/pkg/mod \
	-v $(DIR):/go/src/${PKG}/ \
	-e GO111MODULE=on \
	-e GOOS=linux \
	-e GOARCH=amd64 \
	-w="/go/src/${PKG}/" \
	golang:$(DOCKER_GO_VER) go get -v

## build using docker golang
test-docker:
	docker run \
	-v $(PKGMODPATH):/go/pkg/mod \
	-v $(DIR):/go/src/${PKG}/ \
	-e GO111MODULE=on \
	-e GOOS=linux \
	-e GOARCH=amd64 \
	-w="/go/src/${PKG}/" \
	golang:$(DOCKER_GO_VER) /bin/bash -c " \
	go fmt ./... && \
	go vet ./... && \
	go test -v ./..."

## build using docker golang
build-docker-linux:
	docker build . -t lagoon/lagoon-cli
	docker run -v $(DIR):/app --entrypoint="/bin/sh" lagoon/lagoon-cli -c "cp /lagoon /app/builds/lagoon"

build-docker-darwin:
	docker build . -t lagoon/lagoon-cli-darwin --build-arg OS=darwin --build-arg ARCH=arm64
	docker run -v $(DIR):/app --entrypoint="/bin/sh" lagoon/lagoon-cli-darwin -c "cp /lagoon /app/builds/lagoon"

install-linux:
	cp builds/lagoon-cli-${VERSION}-linux-amd64 ${ARTIFACT_DESTINATION}/lagoon
install-darwin:
	cp builds/lagoon-cli-${VERSION}-darwin-amd64 ${ARTIFACT_DESTINATION}/lagoon

# Settings for the MKDocs serving
MKDOCS_IMAGE ?= ghcr.io/amazeeio/mkdocs-material
MKDOCS_SERVE_PORT ?= 8000

.PHONY: docs/serve
docs/serve:
	@echo "Starting container to serve documentation"
	@docker pull $(MKDOCS_IMAGE)
	@docker run --rm -it \
		-p 127.0.0.1:$(MKDOCS_SERVE_PORT):$(MKDOCS_SERVE_PORT) \
		-v ${PWD}:/docs \
		--entrypoint sh $(MKDOCS_IMAGE) \
		-c 'mkdocs serve -s --dev-addr=0.0.0.0:$(MKDOCS_SERVE_PORT) -f mkdocs.yml'