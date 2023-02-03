ifndef VERBOSE
MAKEFLAGS += --silent
endif

.PHONY: clean
clean:
	rm -rf .cache/*

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	docker pull golangci/golangci-lint:latest > /dev/null \
	&& mkdir -p .cache/golangci-lint \
	&& docker run --rm \
		-v $(shell pwd):/app \
		-v $(shell pwd)/.cache:/root/.cache \
		-w /app golangci/golangci-lint:latest golangci-lint run ./...

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
	docker pull sonatypecommunity/nancy:latest > /dev/null \
	&& go list -buildvcs=false -deps -json ./... \
	| docker run --rm -i sonatypecommunity/nancy:latest sleuth

.PHONY: spell-check
spell-check:
	docker pull ghcr.io/streetsidesoftware/cspell:latest > /dev/null \
	&& docker run --rm \
		-v $(shell pwd):/workdir \
		ghcr.io/streetsidesoftware/cspell:latest \
			--config .vscode/cspell.json "**"

.PHONY: test
test:
	go test -cover ./...

.PHONY: test-json
test-json:
	go test -cover -json ./...

.PHONY: tidy
tidy:
	go mod tidy
