start:
	@echo "==>Starting application..."
	docker-compose up --d apiBank-db
	go run cmd/main.go

stop:
	docker-compose stop apiBank-db

test:
	@echo "==> Running Tests..."
	go test -v ./...

test-coverage:
	@echo "==> Check coverage tests..."
	curl https://gotest-release.s3.amazonaws.com/gotest_linux > gotest && chmod +x gotest
	./gotest -race -failfast -timeout 5m -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	@echo "Running golangci-lint"
ifeq (, $(shell which $$(go env GOPATH)/bin/golangci-lint))
	@echo "==> installing golangci-lint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.44.0
endif
	$$(go env GOPATH)/bin/golangci-lint run --fix --allow-parallel-runners -c ./.golangci.yml ./...

clean:
	@echo "==>Cleanning..."
	rm -f coverage.html
	rm -f coverage.out
	rm -f gotest

