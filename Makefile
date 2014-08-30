default: build

cleanup:
	go clean
	go fmt ./...

build: cleanup
	godep go build

test: cleanup
	godep go test ./...

run:
	godep go run main.go
