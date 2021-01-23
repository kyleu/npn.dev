# Build image
FROM golang:alpine AS builder

ENV GOFLAGS="-mod=readonly"

RUN apk add --update --no-cache bash ca-certificates make git curl build-base

RUN mkdir /npn

WORKDIR /npn

RUN go get -u github.com/pyros2097/go-embed
RUN go get -u github.com/shiyanhui/hero/hero
RUN go get -u golang.org/x/tools/cmd/goimports

ADD ./go.mod          /npn/go.mod
ADD ./go.sum          /npn/go.sum
ADD ./libnpn          /npn/libnpn

RUN go mod download

ADD ./app             /npn/app
ADD ./bin             /npn/bin
ADD ./main.go         /npn/main.go
ADD ./Makefile        /npn/Makefile
ADD ./web             /npn/web

RUN go mod download

RUN set -xe && bash -c 'make build-release-ci'

RUN mv build/release /build

# Final image
FROM alpine

RUN apk add --update --no-cache ca-certificates tzdata bash curl

SHELL ["/bin/bash", "-c"]

COPY --from=builder /build/* /usr/local/bin/

EXPOSE 10101
CMD ["npn"]
