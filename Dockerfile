# Build image
FROM golang:alpine AS builder

ENV GOFLAGS="-mod=readonly"

RUN apk add --update --no-cache ca-certificates make git curl build-base

RUN mkdir /npn

WORKDIR /npn

RUN go get -u github.com/pyros2097/go-embed
RUN go get -u github.com/shiyanhui/hero/hero
RUN go get -u golang.org/x/tools/cmd/goimports

ADD ./.git     /npn/.git
ADD ./Makefile /npn/Makefile
ADD ./go.mod   /npn/go.mod
ADD ./go.sum   /npn/go.sum
ADD ./app      /npn/app
ADD ./bin      /npn/bin
ADD ./client   /npn/client
ADD ./cmd      /npn/cmd
ADD ./web      /npn/web

ARG BUILD_TARGET

COPY go.* /npn/
RUN go mod download

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
