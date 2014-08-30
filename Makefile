default: build

cleanup:
	go clean
	go fmt ./...

build: cleanup
	go build

test: cleanup
	go test ./...

run:
	go run main.go
