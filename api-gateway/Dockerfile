FROM golang:1.18 as builder
#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/jaeger/gateway
WORKDIR $GOPATH/src/gitlab.udevs.io/jaeger/gateway
# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/gateway /

FROM alpine
COPY --from=builder gateway .
RUN apk add --no-cache tzdata
ENV TZ Asia/Tashkent


CMD ["./gateway"]
