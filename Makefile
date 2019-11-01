DIR := $(PWD)
VERSION := latest
GOCMD=go

ARTIFACT_NAME=lagoon
ARTIFACT_DESTINATION=$(GOPATH)/bin

all: deps test build docker-build
deps:
	GO111MODULE="on" ${GOCMD} get -v
test:
	$(GOCMD) fmt ./...
	$(GOCMD) vet ./...
	$(GOCMD) test -v ./...
clean:
	$(GOCMD) clean
build:
	GO111MODULE="on" $(GOCMD) build -o ${ARTIFACT_DESTINATION}/${ARTIFACT_NAME} -v

build-linux:
	GO111MODULE="on" $(GOCMD) build -o builds/lagoon-cli-${VERSION}-linux-amd64 -v
build-darwin:
	GO111MODULE="on" $(GOCMD) build -o builds/lagoon-cli-${VERSION}-darwin-amd64 -v
docker-build:
	docker build -t lagoon-cli -f Dockerfile.build .
	docker run -v $(DIR):/data lagoon-cli cp lagoon-cli-linux-amd64 /data/builds/lagoon-cli-${VERSION}-linux-amd64
	docker run -v $(DIR):/data lagoon-cli cp lagoon-cli-darwin-amd64 /data/builds/lagoon-cli-${VERSION}-darwin-amd64
docker-clean:
	docker image rm lagoon-cli
install-linux:
	cp builds/lagoon-cli-${VERSION}-linux-amd64 ${ARTIFACT_DESTINATION}/lagoon
install-darwin:
	cp builds/lagoon-cli-${VERSION}-darwin-amd64 ${ARTIFACT_DESTINATION}/lagoon