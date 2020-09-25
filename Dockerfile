# Build image
FROM golang:alpine AS builder

ENV GOFLAGS="-mod=readonly"

RUN apk add --update --no-cache bash ca-certificates make git curl build-base

RUN mkdir /npn

WORKDIR /npn

RUN go get -u github.com/pyros2097/go-embed
RUN go get -u github.com/shiyanhui/hero/hero
RUN go get -u golang.org/x/tools/cmd/goimports

ADD ./.git            /npn/.git
ADD ./go.mod          /npn/go.mod
ADD ./go.sum          /npn/go.sum
ADD ./app             /npn/app
ADD ./client          /npn/client
ADD ./cmd             /npn/cmd
ADD ./npnasset        /npn/npnasset
ADD ./npnconnection   /npn/npnconnection
ADD ./npncontroller   /npn/npncontroller
ADD ./npncore         /npn/npncore
ADD ./npndatabase     /npn/npndatabase
ADD ./npnexport       /npn/npnexport
ADD ./npngraphql      /npn/npngraphql
ADD ./npnscript       /npn/npnscript
ADD ./npnservice      /npn/npnservice
ADD ./npnservice-db   /npn/npnservice-db
ADD ./npnservice-fs   /npn/npnservice-fs
ADD ./npntemplate     /npn/npntemplate
ADD ./npnuser         /npn/npnuser
ADD ./npnweb          /npn/npnweb
ADD ./web             /npn/web

ARG BUILD_TARGET

COPY go.* /npn/
RUN go mod download

ADD ./bin             /npn/bin
ADD ./Makefile        /npn/Makefile

RUN set -xe && make build-release-force

RUN mv build/release /build

# Final image
FROM alpine

RUN apk add --update --no-cache ca-certificates tzdata bash curl

SHELL ["/bin/bash", "-c"]

# set up nsswitch.conf for Go's "netgo" implementation
# https://github.com/gliderlabs/docker-alpine/issues/367#issuecomment-424546457
RUN test ! -e /etc/nsswitch.conf && echo 'hosts: files dns' > /etc/nsswitch.conf

ARG BUILD_TARGET

RUN if [[ "${BUILD_TARGET}" == "debug" ]]; then apk add --update --no-cache libc6-compat; fi

COPY --from=builder /build/* /usr/local/bin/

EXPOSE 10101
CMD ["npn"]
