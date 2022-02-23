.PHONY: build
build:
	@echo 'Start build'
	@go build -o bin/crawler cmd/crawler/main.go
	@echo 'The app was successfully built at ./bin/crawler '

.PHONY: all
all: lint test build run

.PHONY: run
run: build
	@echo 'Start app'
	@./bin/crawler -h

.PHONY: test
test:
	@echo 'Start tests'
	go test -race -coverprofile=coverage.txt -covermode=atomic --tags=integration ./...

.PHONY: lint
lint:
	@echo 'Start lint'
	golangci-lint run

.PHONY: clean
clean:
	@echo '>> cleaning go'
	@go clean
	@echo '>> cleaning binaries'
	@-rm -rf bin
