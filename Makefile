include .env
export

build:
	go build -o bin/app ./cmd/main.go
fmt:
	go fmt ./...
tidy:
	go mod tidy
lint:
	go vet ./...
test:
	go test ./...
clean:
	rm -rf bin
start:
	./bin/app