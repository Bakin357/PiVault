BINARY_NAME = PiVault

build:
	@echo "Building the project..."
	go build -o bin/$(BINARY_NAME) ./src/cmd/main.go

run: build
	@echo "Running the project..."
	./bin/$(BINARY_NAME)

test: 
	@echo "Running tests..."
	go test ./...

clean:
	@echo "Cleaning up..."
	rm -f bin/$(BINARY_NAME)

.PHONY: build run test clean