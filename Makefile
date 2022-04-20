.PHONY: test

all: build

build:
	@go build -v .

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf task-manager
	go clean -i .

test:
	sh -c 'go test -count=1 ./...'