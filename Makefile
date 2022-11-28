ifndef VERBOSE
MAKEFLAGS += --silent
endif

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: mock
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

.PHONY: nancy
nancy:
	go list -json -m all | nancy sleuth

.PHONY: spell-check
spell-check:
	# npm install -g cspell@latest
	cspell lint --config .vscode/cspell.json ".*" && \
	cspell lint --config .vscode/cspell.json "**/.*" && \
	cspell lint --config .vscode/cspell.json ".{github,vscode}/**/*" && \
	cspell lint --config .vscode/cspell.json "**"

.PHONY: test
test:
	go test -cover ./...

.PHONY: test-json
test-json:
	go test -cover -json ./...

.PHONY: tidy
tidy:
	go mod tidy
