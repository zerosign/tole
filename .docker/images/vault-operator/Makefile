REGISTRY ?= docker.pkg.github.com
ORG      ?= zerosign
PROJECT  ?= tole/vault-operator
VERSION  ?= $(strip $(shell git show -q --format=%h))


all: build

build:
	docker build -f Dockerfile -t $(REGISTRY)/$(ORG)/$(PROJECT):$(VERSION) ../../../ --no-cache
