BINARY_NAME = broLang

.PONY: build
build: clean
	@go build -o bin/$(BINARY_NAME) 
	@if [ $$? -eq 0 ]; then \
		echo "\033[32mBuild Success\033[0m"; \
	else \
		echo "\033[31mBuild Failed\033[0m"; \
	fi

.PHONY: test
test:
	@go test -v ./test

.PHONY: clean
clean:
	@go clean
	@rm -f bin/$(BINARY_NAME)