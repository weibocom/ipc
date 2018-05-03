WORKDIR=`pwd`

default: build

vet:
	go vet ./...

tools:
	go get honnef.co/go/tools/cmd/staticcheck
	go get honnef.co/go/tools/cmd/gosimple
	go get honnef.co/go/tools/cmd/unused
	go get github.com/gordonklaus/ineffassign
	go get github.com/fzipp/gocyclo
	go get github.com/golang/lint/golint
	go get github.com/alexkohler/prealloc

gometalinter:
	gometalinter --enable-all ./...
lint:
	golint ./...

staticcheck:
	staticcheck -ignore "$(shell cat .checkignore)" ./...

gosimple:
	gosimple -ignore "$(shell cat .gosimpleignore)" ./...

unused:
	unused ./...

ineffassign:
	ineffassign .

gocyclo:
	gocyclo -over 20 $(shell find . -name "*.go" |egrep -v "_testutils/*|vendor/*|pb\.go|_test\.go")

prealloc:
	prealloc ./...

check: staticcheck gosimple ineffassign

fmt:
	go fmt ./...

build:
	go build ./...

test:
	go test -race -v ./...
