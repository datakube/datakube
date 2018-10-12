#docker build --rm -t datakube/server .
FROM node:9.11.1-alpine as build-webui
WORKDIR /app
COPY webui/package*.json ./
RUN yarn install
COPY webui/ .
RUN yarn run build

####
FROM golang:1.11-alpine as build-server

ARG DEP_VERSION=0.4.1

RUN apk --update upgrade \
    && apk --no-cache --no-progress add git bash gcc musl-dev curl tar make\
    && rm -rf /var/cache/apk/*

RUN mkdir -p /usr/local/bin \
    && curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 \
    && chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/datakube/datakube
COPY . /go/src/github.com/datakube/datakube
COPY --from=build-webui /app/dist /go/src/github.com/datakube/datakube/dist
RUN go get github.com/rakyll/statik && \
    go generate && \
    CGO_ENABLED=0 GOGC=off go build -o dist/server ./cmd/server/

#####
FROM alpine:3.6

EXPOSE 8080
ENV GIN_MODE=release

COPY --from=build-server /go/src/github.com/datakube/datakube/dist/server /bin/datakube
ENTRYPOINT ["/bin/datakube"]

