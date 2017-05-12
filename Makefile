SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
WORKER_MOUNT := -v "$(CURDIR)/$(BIND_DIR):/go/src/github.com/SantoDE/datahamster/$(BIND_DIR)"
BIND_DIR := "dist"

GIT_BRANCH := $(subst heads/,,$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null))

default: binary

binary:
	./script/make.sh binary

crossbinary:
	./script/make.sh crossbinary

image:
	docker build -t $(WORKER_IMAGE) .

lint:
	./script/make.sh validate-golint

fmt:
	gofmt -s -l -w $(SRCS)

validate:  ## validate gofmt, golint and go vet
	./script/make.sh validate-gofmt validate-golint

test-unit:
	./script/make.sh test-unit

test-integration:
	 ./script/make.sh test-integration

integration-test-image:
	docker build -f integration/Dockerfile.Integration -t santode/datahamster-worker-integration-test-db .