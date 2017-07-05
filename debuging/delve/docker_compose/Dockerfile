FROM golang

RUN mkdir -p /opt/src

ENV GOPATH "$GOPATH:/opt"

WORKDIR /opt/src/mocking

# Install extra networking packages
RUN go get github.com/derekparker/delve/cmd/dlv

