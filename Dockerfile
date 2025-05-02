FROM golang:1.24-alpine3.20 AS build

WORKDIR /go/src/github.com/uselagoon/lagoon-cli/
COPY . .

ARG VERSION
ARG OS
ENV OS=${OS:-linux}
ARG ARCH
ENV ARCH=${ARCH:-amd64}

RUN apk update && apk add git

RUN VERSION=${VERSION:-"$(echo $(git describe --abbrev=0 --tags)+$(git rev-parse --short=8 HEAD))"} \
	&& BUILD=$(date +%FT%T%z) \
  	&& CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build \
	-ldflags "-w -s -X github.com/uselagoon/lagoon-cli/cmd.lagoonCLIVersion=$VERSION \
	-X github.com/uselagoon/lagoon-cli/cmd.lagoonCLIBuild=$BUILD \
	-X github.com/uselagoon/lagoon-cli/cmd.lagoonCLIBuildGoVersion=go$GOLANG_VERSION" -o lagoon .

FROM alpine:3.21

WORKDIR /root/
COPY --from=build /go/src/github.com/uselagoon/lagoon-cli/lagoon /lagoon

RUN touch ~/.lagoon.yml

ENTRYPOINT ["/lagoon"]
