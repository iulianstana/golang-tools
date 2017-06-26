FROM golang

RUN mkdir -p /opt/src

ENV GOPATH "$GOPATH:/opt"

WORKDIR /opt/src/mocking

# Install extra networking packages
RUN go get -u github.com/go-redis/redis
RUN go get gopkg.in/redis.v3
RUN go get github.com/vektra/mockery/.../
