GOCMD=go

ARTIFACT_NAME=lagoon
ARTIFACT_DESTINATION=$(GOPATH)/bin

all: deps test build
deps:
	GO111MODULE="on" ${GOCMD} get -v
build:
	GO111MODULE="on" $(GOCMD) build -o $(ARTIFACT_DESTINATION)/$(ARTIFACT_NAME) -v
test:
	$(GOCMD) fmt ./...
	$(GOCMD) vet ./...
	$(GOCMD) test -v ./...
clean:
	$(GOCMD) clean
