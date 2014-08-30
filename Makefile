clean:
	go fmt

build: clean
	go build

test: clean
	go test ./...
