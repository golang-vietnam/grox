default: build test

install:
	godep restore

build:
	go install github.com/golang-vietnam/grox/cmd/grox-server

test:
	go test -v ./...

run:
	$(GOPATH)/bin/grox-server

start: build run

