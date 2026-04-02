ifndef VERBOSE
MAKEFLAGS += --silent
endif

pkg ?= ./...
pwd = $(shell pwd)

.PHONY: ci-suite
ci-suite: spell-check fmt lint lint-markdown mock vulnerability-check test

.PHONY: clean
clean:
	rm -rf .cache/*

.PHONY: fmt
fmt:
	go fmt $(pkg)

.PHONY: lint
lint:
	mkdir -p .cache/golangci-lint .cache/go-build \
	&& docker run --pull always --rm \
		-v "$(pwd):/app" \
		-v "$(pwd)/.cache:/root/.cache" \
		-w /app \
		golangci/golangci-lint:latest golangci-lint run $(pkg)

.PHONY: lint-markdown
lint-markdown:
	docker run --pull always --rm -v "$(pwd):$(pwd)" -w "$(pwd)" \
		davidanson/markdownlint-cli2:latest "**/*.md" \
		--config .vscode/.markdownlint-cli2.yaml

.PHONY: mock
mock:
	docker run --pull always --rm -v "$(pwd):/src" -w /src vektra/mockery

.PHONY: spell-check
spell-check:
	docker run --pull always --rm -v "$(pwd):/workdir" \
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

.PHONY: vulnerability-check
vulnerability-check:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck -show=version ./...
