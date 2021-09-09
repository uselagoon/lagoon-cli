DIR := $(PWD)
GOCMD=go

ARTIFACT_NAME=lagoon
ARTIFACT_DESTINATION=$(GOPATH)/bin

PKG=github.com/uselagoon/lagoon-cli
PKGMODPATH=$(DIR)/vendor

VERSION=$(shell ${PWD}/increment_ver.sh -p $(shell git describe --abbrev=0 --tags))-rc
BUILD=$(shell date +%FT%T%z)

DOCKER_GO_VER=1.14
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

docs: test build
	GO111MODULE=on $(GOCMD) run main.go --docs

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
	docker run \
	-v $(PKGMODPATH):/go/pkg/mod \
	-v $(DIR):/go/src/${PKG}/ \
	-e GO111MODULE=on \
	-e GOOS=linux \
	-e GOARCH=amd64 \
	-w="/go/src/${PKG}/" \
	golang:$(DOCKER_GO_VER) go build -ldflags '${LDFLAGS} -X "${PKG}/cmd.lagoonCLIBuildGoVersion=${GO_VER}"' -o builds/lagoon-cli-${VERSION}-linux-amd64

build-docker-darwin:
	docker run \
	-v $(PKGMODPATH):/go/pkg/mod \
	-v $(DIR):/go/src/${PKG}/ \
	-e GO111MODULE=on \
	-e GOOS=darwin \
	-e GOARCH=amd64 \
	-w="/go/src/${PKG}/" \
	golang:$(DOCKER_GO_VER) go build -ldflags '${LDFLAGS} -X "${PKG}/cmd.lagoonCLIBuildGoVersion=${GO_VER}"' -o builds/lagoon-cli-${VERSION}-darwin-amd64

install-linux:
	cp builds/lagoon-cli-${VERSION}-linux-amd64 ${ARTIFACT_DESTINATION}/lagoon
install-darwin:
	cp builds/lagoon-cli-${VERSION}-darwin-amd64 ${ARTIFACT_DESTINATION}/lagoon

release-patch: 
	$(eval VERSION=$(shell ${PWD}/increment_ver.sh -p $(shell git describe --abbrev=0 --tags)))
	git tag $(VERSION)
	mkdocs gh-deploy
	git push $(GIT_ORIGIN) main --tags

release-minor: 
	$(eval VERSION=$(shell ${PWD}/increment_ver.sh -m $(shell git describe --abbrev=0 --tags)))
	git tag $(VERSION)
	mkdocs gh-deploy
	git push $(GIT_ORIGIN) main --tags

release-major: 
	$(eval VERSION=$(shell ${PWD}/increment_ver.sh -M $(shell git describe --abbrev=0 --tags)))
	git tag $(VERSION)
	mkdocs gh-deploy
	git push $(GIT_ORIGIN) main --tags
