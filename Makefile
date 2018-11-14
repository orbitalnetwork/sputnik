# GOPATH := $(shell pwd)/../../..
CURRENT := $(shell pwd)

VERSION=`git describe --tags --always`
VERSION_DEV = `git log -1 --pretty=format:"%H"`
BUILD=`date +%F+%T%z`

# enable scenario supportr (JavaScript)
GO_TAGS += scenario

# enable control interface (development porposes)
# GO_TAGS += control
all: sputnik

install:
	@GOPATH=$(GOPATH) go install -tags "$(GO_TAGS)" -ldflags " -s -X main.Version=${VERSION} -X main.Build=${BUILD} "

sputnik:
	mkdir -p $(CURRENT)/build
	@GOPATH=$(GOPATH) go build -tags "$(GO_TAGS)" -ldflags " -X main.Version=${VERSION_DEV} -X main.Build=${BUILD} " -o build/$@ cmd/sputnik.go

clean:
	@GOPATH=$(GOPATH) go clean

test:
	GOCACHE=off go test --race ./pkg/...

dep:
	dep ensure

.PHONY: all install test sputnik clean

