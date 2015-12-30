.PHONY: build shell test-unit

PROJECT=filewatcher
BUILD_ID ?= $(shell git rev-parse --short HEAD 2>/dev/null)
DOCKER_IMAGE := $(PROJECT)-dev:$(BUILD_ID)

VOLUMES := \
	-v $(CURDIR):/go/src/github.com/dnephin/$(PROJECT) \
	-v $(CURDIR)/dist/bin:/go/bin \
	-v $(CURDIR)/dist/pkg:/go/pkg

all: binary

build:
	docker build -t $(DOCKER_IMAGE) -f Dockerfile.build .

shell: build
	docker run --rm -ti $(VOLUMES) $(DOCKER_IMAGE) bash

test-unit: build
	docker run --rm -ti $(VOLUMES) $(DOCKER_IMAGE) go test -v ./...

test: test-unit

binary: build
	docker run --rm -ti $(VOLUMES) $(DOCKER_IMAGE) go install