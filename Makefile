APP_NAME=spotify-cli

ifeq ($(OS),Windows_NT)
    EXE_EXT := .exe
    RUN_CMD := .\$(APP_NAME)$(EXE_EXT)
else
    EXE_EXT :=
    RUN_CMD := ./$(APP_NAME)p$(EXE_EXT)
endif

#################################################################
# Pre-configuration section
#################################################################

.PHONY: install-deps
install-deps:
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/vektra/mockery/v2@latest

.PHONY: tidy
tidy:
	go mod tidy

#################################################################
# Build and bun section
#################################################################

.PHONY: build
build:
	go build -o ./$(APP_NAME)$(EXE_EXT) ./cmd/$(APP_NAME) || { echo "build failed"; exit 1; }


.PHONY: run
run: build
	$(RUN_CMD)


#################################################################
# Lint section
#################################################################

.PHONY: lint
lint:
	golangci-lint run || { echo "lint failed"; exit 1; }

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix  || { echo "lint-fix failed"; exit 1; }


#################################################################
# Tests and mocks section
#################################################################

.PHONY: test
test:
	go test ./... || { echo "tests failed"; exit 1; }


.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage on internal package"
	$(eval PKGS := $(shell go list ./internal/... | grep -vE '(/model(/|$$)|/mocks(/|$$))'))

	@if [ -z "$(PKGS)" ]; then \
		echo "No packages found to test after filtering."; \
		exit 1; \
	fi
	go test -coverprofile=coverage.out $(PKGS)

.PHONY: test-coverage-detailed
test-coverage-detailed: test-coverage
	@echo "Generating detailed coverage report to coverage_detailed.txt..."
	go tool cover -func=coverage.out > coverage_detailed.txt
    @echo "Detailed coverage report generated: coverage_detailed.txt"


.PHONY: test-coverage-html
test-coverage-html: test-coverage
	@echo "Generating HTML coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "HTML coverage report generated: coverage.html"


.PHONY: mocks-gen
mocks-gen:
	mockery --all --recursive --dir ./internal --output ./internal/mocks --keeptree --case snake


#################################################################
# Other
#################################################################

.PHONY: clean
clean:
	@echo "Cleaning up test coverage files..."
	rm -f coverage.out coverage.html coverage_detailed.txt spotify-cli.exe


#################################################################
# Developer sections
#################################################################

.PHONY: pre-commit
pre-commit: mocks-gen lint-fix build test test-coverage-detailed clean
