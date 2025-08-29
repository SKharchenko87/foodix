build:
	go build -o ./bin/foodix ./cmd/server/main.go
run: build
	./bin/foodix
test:
	go test ./... -cover
lint:
	golangci-lint run ./...
clean:
	go clean
	rm -f ./bin/foodix
vet:
	go vet
tidy:
	go mod tidy
swag:
	swag init --dir ./cmd/server/,./