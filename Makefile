.PHONY: all deps test build clean

REPO=github.com/svenwltr/json-flatten

VERSION=$(shell git describe --always --dirty | tr '-' '.' )

all: test build cov

bootstrap: tools deps

clean:
	rm -rf target vendor

deps:
ifneq ($(strip $(shell go list)),$(REPO))
	$(error This repository should be cloned into $$GOPATH/src/$(REPO))
endif
	glide install

test: deps
	go test -cover $(shell glide nv)

build: deps
	mkdir -p target
	go build \
		-ldflags "-X main.version=$(VERSION)" \
		-o target/json-flatten

tools:
	go get gopkg.in/matm/v1/gocov-html
	go get github.com/axw/gocov/gocov

cov: deps tools
	gocov test $(shell glide nv) \
		| gocov-html > target/coverage.html
