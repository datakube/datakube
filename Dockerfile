#docker build --rm -t datahamster/server .

FROM golang:1.11-alpine as builder

ARG DEP_VERSION=0.4.1

RUN apk --update upgrade \
    && apk --no-cache --no-progress add git bash gcc musl-dev curl tar \
    && rm -rf /var/cache/apk/*

RUN mkdir -p /usr/local/bin \
    && curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 \
    && chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/SantoDE/datahamster
COPY . /go/src/github.com/SantoDE/datahamster

RUN CGO_ENABLED=0 GOGC=off go build $FLAGS -ldflags "-X github.com/SantoDE/datahamster/cmd/server/version.Version=$VERSION -X github.com/SantoDE/datahamster/cmd/server/version.BuildDate=$DATE" -o dist/server ./cmd/server/

FROM alpine:3.6

COPY --from=builder /go/src/github.com/SantoDE/datahamster/dist/server /bin/datahamster
ENTRYPOINT ["/bin/datahamster"]

