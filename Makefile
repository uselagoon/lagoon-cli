GOCMD=go
DEPCMD=dep

ARTIFACT_NAME=lagoon
ARTIFACT_DESTINATION=$(GOPATH)/bin

all: dep test build
deps:
	$(DEPCMD) ensure
build:
	$(GOCMD) build -o $(ARTIFACT_DESTINATION)/$(ARTIFACT_NAME) -v
test:
	$(GOCMD) fmt ./...
	$(GOCMD) vet ./...
	$(GOCMD) test -v ./...
clean:
	$(GOCMD) clean
