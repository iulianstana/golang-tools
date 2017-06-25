FROM golang

RUN mkdir -p /opt/src

ENV GOPATH "$GOPATH:/opt"

WORKDIR /opt/src/gore-shell

# Install extra networking packages
RUN go get -u github.com/motemen/gore
