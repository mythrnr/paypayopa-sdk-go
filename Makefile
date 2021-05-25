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

	mockery \
		--case=underscore \
		--dir=$(shell go env GOROOT)/src/net/http \
		--name=RoundTripper \
		--output=internal/mocks

	mockery \
		--case=underscore \
		--dir=$(shell go env GOROOT)/src/io \
		--name=ReadCloser \
		--output=internal/mocks

	mockery \
		--case=underscore \
		--dir=$(shell go env GOROOT)/src/encoding/json \
		--name=Marshaler \
		--output=internal/mocks

test:
	go test -cover ./...

test-json:
	go test -cover -json ./...

tidy:
	go mod tidy
