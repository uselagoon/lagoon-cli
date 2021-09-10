FROM golang:1.14-alpine3.13 as build

WORKDIR /go/src/github.com/uselagoon/lagoon-cli/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags '-w -s -X /go/src/source/cmd.lagoonCLIBuild=1 \
    -X /go/src/source/cmd.lagoonCLIBuildGoVersion=1.14"' \
    -o lagoon .

FROM alpine:3.13 

WORKDIR /root/
COPY --from=build /go/src/github.com/uselagoon/lagoon-cli/lagoon /lagoon
ENTRYPOINT ["/lagoon"]
