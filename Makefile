CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-repo/
	cp *.go src/github.com/whosonfirst/go-whosonfirst-repo/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:   rmdeps
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-cli"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-placetypes"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/wof-repo-parse cmd/wof-repo-parse.go

fmt:
	go fmt *.go
	go fmt cmd/*.go
