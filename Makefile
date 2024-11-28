ifndef VERBOSE
MAKEFLAGS += --silent
endif

pkg ?= ./...
pwd = $(shell pwd)

.PHONY: clean
clean:
	rm -rf .cache/*

.PHONY: fmt
fmt:
	go fmt $(pkg)

.PHONY: lint
lint:
	docker pull golangci/golangci-lint:latest > /dev/null \
	&& mkdir -p .cache/golangci-lint .cache/go-build \
	&& docker run --rm \
		-v $(pwd):/app \
		-v $(pwd)/.cache:/root/.cache \
		-w /app \
		golangci/golangci-lint:latest golangci-lint run $(pkg)

.PHONY: mock
mock:
	docker pull vektra/mockery:latest > /dev/null

	docker run --rm \
		-v $(pwd):/src \
		-w /src vektra/mockery \
		--case=underscore \
		--dir=/usr/local/go/src/net/http \
		--name=RoundTripper \
		--output=internal/mocks

	docker run --rm \
		-v $(pwd):/src \
		-w /src vektra/mockery \
		--case=underscore \
		--dir=/usr/local/go/src/io \
		--name=ReadCloser \
		--output=internal/mocks

	docker run --rm \
		-v $(pwd):/src \
		-w /src vektra/mockery \
		--case=underscore \
		--dir=/usr/local/go/src/encoding/json \
		--name=Marshaler \
		--output=internal/mocks

.PHONY: nancy
nancy:
	docker pull sonatypecommunity/nancy:latest > /dev/null \
	&& go list -buildvcs=false -deps -json $(pkg) \
	| docker run --rm -i sonatypecommunity/nancy:latest sleuth

.PHONY: spell-check
spell-check:
	docker pull ghcr.io/streetsidesoftware/cspell:latest > /dev/null \
	&& docker run --rm \
		-v $(pwd):/workdir \
		ghcr.io/streetsidesoftware/cspell:latest \
			--config /workdir/.vscode/cspell.json "**"

.PHONY: release
release:
	if [ "$(version)" = "" ]; then \
		echo "version is required."; \
		exit 1; \
	fi \
	&& gh release create $(version) --generate-notes --target master

.PHONY: test
test:
	go test -cover $(pkg)

.PHONY: test-json
test-json:
	go test -cover -json $(pkg)

.PHONY: tidy
tidy:
	go mod tidy
