WORKER_IMAGE := $(if $(REPONAME),$(REPONAME),"santode/datahamster-worker")

default: binary

binary:
	./script/make.sh binary

crossbinary:
	./script/make.sh crossbinary

image:
	docker build -t $(WORKER_IMAGE) .