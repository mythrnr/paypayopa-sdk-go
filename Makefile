.PHONY: lint mock test test-json tidy
.SILENT: test-json

lint:
	golangci-lint run ./...

mock:
	mockery \
		--case=underscore \
		--inpackage \
		--name=.* \
		--output .

test:
	go test -cover ./...

test-json:
	go test -cover -json ./...

tidy:
	go mod tidy
