DIR := $(PWD)
GOCMD=go

ARTIFACT_NAME=lagoon
ARTIFACT_DESTINATION=$(GOPATH)/bin

PKG=github.com/amazeeio/lagoon-cli
PKGMODPATH=$(DIR)/vendor

VERSION=$(shell ${PWD}/increment_ver.sh -p $(shell git describe --abbrev=0 --tags))-rc
BUILD=$(shell date +%FT%T%z)

DOCKER_GO_VER=1.14
GO_VER=$(shell go version)
LDFLAGS=-w -s -X ${PKG}/cmd.lagoonCLIVersion=${VERSION} -X ${PKG}/cmd.lagoonCLIBuild=${BUILD}


all: deps test build docs
all-docker-linux: deps-docker test-docker build-docker-linux
all-docker-darwin: deps-docker test-docker build-docker-darwin

deps:
	GO111MODULE=on ${GOCMD} get -v
test:
	GO111MODULE=on $(GOCMD) fmt ./...
	GO111MODULE=on $(GOCMD) vet ./...
	GO111MODULE=on $(GOCMD) test -v ./...
gen:
	GO111MODULE=on $(GOCMD) generate ./...

clean:
	$(GOCMD) clean

build:
	GO111MODULE=on $(GOCMD) build -ldflags '${LDFLAGS} -X "${PKG}/cmd.lagoonCLIBuildGoVersion=${GO_VER}"' -o ${ARTIFACT_DESTINATION}/${ARTIFACT_NAME} -v
build-linux:
	GO111MODULE=on GOOS=linux GOARCH=amd64 $(GOCMD) build -ldflags '${LDFLAGS} -X "${PKG}/cmd.lagoonCLIBuildGoVersion=${GO_VER}"' -o builds/lagoon-cli-${VERSION}-linux-amd64 -v
build-darwin:
	GO111MODULE=on GOOS=darwin GOARCH=amd64 $(GOCMD) build -ldflags '${LDFLAGS} -X "${PKG}/cmd.lagoonCLIBuildGoVersion=${GO_VER}"' -o builds/lagoon-cli-${VERSION}-darwin-amd64 -v

docs: test
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
	git push origin master --tags

release-minor: 
	$(eval VERSION=$(shell ${PWD}/increment_ver.sh -m $(shell git describe --abbrev=0 --tags)))
	git tag $(VERSION)
	mkdocs gh-deploy
	git push origin master --tags

release-major: 
	$(eval VERSION=$(shell ${PWD}/increment_ver.sh -M $(shell git describe --abbrev=0 --tags)))
	git tag $(VERSION)
	mkdocs gh-deploy
	git push origin master --tags
