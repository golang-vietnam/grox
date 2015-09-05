default: build test

install:
	godep restore

build:
	go install github.com/golang-vietnam/grox/cmd/grox-server

save:
	@echo "Saving..."
	@go get github.com/kr/godep	
	@export PATH=$(PATH):$(GOPATH)/bin;godep save
	@echo "Save new dependencies successful!"

test:
	go test -v ./...

start:
	$(GOPATH)/bin/grox-server

run: build start

