SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
WORKER_MOUNT := -v "$(CURDIR)/$(BIND_DIR):/go/src/github.com/datakube/datakube/$(BIND_DIR)"
BIND_DIR := "dist"

GIT_BRANCH := $(subst heads/,,$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null))

default: binary

images:
	docker build -t datakube/server .
	docker build -f Dockerfile.agent -t datakube/agent .

binary:
	./script/make.sh binary

crossbinary:
	./script/make.sh crossbinary

lint:
	./script/make.sh validate-golint

fmt:
	gofmt -s -l -w $(SRCS)

validate:  ## validate gofmt, golint and go vet
	./script/make.sh validate-gofmt validate-golint

test:
	go test -race ./...

