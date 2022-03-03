dev-local:
	@echo "==>Starting application..."
	docker-compose up --d apiBank-db
	go run cmd/main.go

test:
	echo "==> Running Tests..."
	go test -v ./...

test-coverage:
	echo "==> Check coverage tests..."
	go install github.com/kyoh86/richgo@latest
	@richgo test -failfast -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

cover:
	echo "==> Check coverage tests with go tools"
	go test -cover ./...
