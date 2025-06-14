all: build test

# Build the application
.PHONY: build
build:
	@echo "Building..."
	@go build -o build/main cmd/app/main.go

# Run the application
.PHONY: run
run:
	@go run cmd/app/main.go

.PHONY: start-mongo
start-mongo:
	@docker-compose -f docker-compose.mongo.yaml up -d

.PHONY: start
start:
	@docker-compose up -d

.PHONY: stop
stop:
	@docker-compose down

# Test the application
.PHONY: test
test:
	@echo "Testing..."
	@go test ./... -v

# Live Reload
.PHONY: watch
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

