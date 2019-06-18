REGISTRY ?= docker.pkg.github.com
ORG      ?= zerosign
PROJECT  ?= tole
# VERSION  ?= $(strip $(shell git describe --tags))
VERSION  ?= $(strip $(shell git show -q --format=%h))

.PHONY: clean build test

all: compile test build doc

clean:
	go clean

compile:
	go build

test:
	go test

build:
	docker build -t $(REGISTRY)/$(ORG)/repo/$(PROJECT):$(VERSION) .

push:
	docker push $(REGISTRY)/$(ORG)/repo/$(PROJECT):$(VERSION)

doc:
	go doc
