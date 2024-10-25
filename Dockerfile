FROM golang:1.23-alpine AS build

WORKDIR /go/src/github.com/uselagoon/lagoon-cli/
COPY . .

ARG VERSION

RUN apk update && apk add git

RUN VERSION=${VERSION:-"$(echo $(git describe --abbrev=0 --tags)+$(git rev-parse --short=8 HEAD))"} \
	&& BUILD=$(date +%FT%T%z) \
  && CGO_ENABLED=0 GOOS=linux go build \
	-ldflags "-w -s -X github.com/uselagoon/lagoon-cli/cmd.lagoonCLIVersion=$VERSION \
	-X github.com/uselagoon/lagoon-cli/cmd.lagoonCLIBuild=$BUILD \
	-X github.com/uselagoon/lagoon-cli/cmd.lagoonCLIBuildGoVersion=go$GOLANG_VERSION" -o lagoon .

FROM alpine:3

WORKDIR /root/
COPY --from=build /go/src/github.com/uselagoon/lagoon-cli/lagoon /lagoon

RUN touch ~/.lagoon.yml

ENTRYPOINT ["/lagoon"]
