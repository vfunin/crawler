.PHONY: build
build:
	go build -o bin/crawler cmd/crawler/main.go

.PHONY: all
all:
	@go build -o bin/crawler cmd/crawler/main.go
	@echo 'The app was successfully built at ./bin/crawler '
	@./bin/crawler -h

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test -covermode=atomic ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	@echo '>> cleaning go'
	@go clean
	@echo '>> cleaning binaries'
	@-rm -rf bin
