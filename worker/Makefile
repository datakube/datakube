WORKER_IMAGE := $(if $(REPONAME),$(REPONAME),"santode/datahamster-worker")
SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

default: binary

binary:
	./script/make.sh binary

crossbinary:
	./script/make.sh crossbinary

image:
	docker build -t $(WORKER_IMAGE) .

lint:
	find . -type d -not -path "./vendor/*" | xargs -L 1 golint

fmt:
	gofmt -s -l -w $(SRCS)

test: test-unit test-integration

test-unit:
	go test -test.short ./dumper/sql

test-integration: integration-test-image
	go test ./dumper/sql/

integration-test-image:
	docker build -f Dockerfile.Integration -t santode/datahamster-worker-integration-test-db .