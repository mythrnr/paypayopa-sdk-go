.PHONY: lint
.SILENT: lint
lint:
	golangci-lint run ./...

.PHONY: mock
.SILENT: mock
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

.PHONY: test
.SILENT: test
test:
	go test -cover ./...

.PHONY: test-json
.SILENT: test-json
test-json:
	go test -cover -json ./...

.PHONY: tidy
.SILENT: tidy
tidy:
	go mod tidy
