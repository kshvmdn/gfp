name := gfp
sources := $(wildcard **/*.go)

.PHONY: clean install lint test

all: build

build: $(sources)
	go build -o ./$(name) -v ./cmd/gfp

clean:
	rm -rf ./$(name) build/

install:
	go get -u -v ./...

lint:
	${GOPATH}/bin/golint ./...

test:
	go test ./...
