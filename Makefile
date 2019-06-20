REGISTRY     ?= docker.pkg.github.com
ORG          ?= zerosign
PROJECT      ?= tole
CURRENT_DIR	 := $(shell pwd)
GO_EXISTS    := $(shell command -v go 2> /dev/null)
# VERSION    ?= $(strip $(shell git describe --tags))
VERSION      ?= $(strip $(shell git show -q --format=%h))

.PHONY: clean build test doc clean-compose tools build-compose

all: compile test build doc tools

clean:
	go clean

tools:
	cd .docker/images/vault-operator && make

compile:
ifndef GO_EXISTS
	docker container run  -v $(CURRENT_DIR):"/builder/$(PROJECT)" -w "/builder/$(PROJECT)" -it golang:latest go build
else
	go build
endif
test:
ifndef GO_EXISTS
	docker container run  -v $(CURRENT_DIR):"/builder/$(PROJECT)" -w "/builder/$(PROJECT)" -it golang:latest go vet
	docker container run  -v $(CURRENT_DIR):"/builder/$(PROJECT)" -w "/builder/$(PROJECT)" -it golang:latest go test
else
	go vet
	go test
endif

build:
	docker build -t $(REGISTRY)/$(ORG)/repo/$(PROJECT):$(VERSION) .

push:
	docker push $(REGISTRY)/$(ORG)/repo/$(PROJECT):$(VERSION)

release:
	docker tag $(REGISTRY)/$(ORG)/repo/$(PROJECT):latest
	docker push $(REGISTRY)/$(ORG)/repo/$(PROJECT):latest

build-compose:
	docker-compose up --no-start

clean-compose:
	docker-compose down -v

doc:
ifndef GO_EXISTS
	docker container run  -v $(CURRENT_DIR):"/builder/$(PROJECT)" -w "/builder/$(PROJECT)" -it golang:latest go doc
else
	go doc
endif
