.PHONY: test

# Test the handler
test:
	go test ./... -v

# Run the application
run:
	go run cmd/main.go