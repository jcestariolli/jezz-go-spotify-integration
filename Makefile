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
	@go mod download
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/vektra/mockery/v2@latest

.PHONY: tidy
tidy:
	@go mod tidy

#################################################################
# Build and bun section
#################################################################

.PHONY: build
build:
	@echo "####### Building project..."
	@go build -o ./$(APP_NAME)$(EXE_EXT) ./cmd/$(APP_NAME) || { echo "####### ERROR: Build failed"; exit 1; }
	@echo "####### Building succeeded!"


.PHONY: run
run: build
	@$(RUN_CMD)


#################################################################
# Lint section
#################################################################

.PHONY: lint
lint:
	@echo "####### Running lint on files..."
	@golangci-lint run || { echo "####### ERROR: Lint failed"; exit 1; }
	@echo "####### Linting succeeded!"

.PHONY: lint-fix
lint-fix:
	@echo "####### Running lint with fix on files..."
	@golangci-lint run --fix  || { echo "####### ERROR: Lint with fix failed"; exit 1; }
	@echo "####### Linting with fix succeeded!"


#################################################################
# Tests and mocks section
#################################################################

.PHONY: test
test:
	@echo "####### Running tests..."
	@go test ./... || { echo "####### ERROR: Tests failed"; exit 1; }
	@echo "####### Tests succeeded!"


.PHONY: test-coverage
test-coverage:
	@echo "####### Running tests with coverage on internal package..."
	@$(eval PKGS := $(shell go list ./internal/... | grep -vE '(/model(/|$$)|/mocks(/|$$))'))

	@if [ -z "$(PKGS)" ]; then \
		echo "####### ERROR: No packages found to test after filtering."; \
		exit 1; \
	fi
	@go test -coverprofile=coverage.out $(PKGS) || { echo "####### ERROR: Tests coverage failed"; exit 1; }
	@echo "####### Tests with coverage succeeded!"


.PHONY: test-coverage-detailed
test-coverage-detailed: test-coverage
	@echo "####### Generating detailed coverage report..."
	@go tool cover -func=coverage.out > coverage_detailed.txt
	@echo "####### Detailed coverage report generated: coverage_detailed.txt"


.PHONY: test-coverage-html
test-coverage-html: test-coverage
	@echo "####### Generating HTML coverage report..."
	@go tool cover -html=coverage.out -o coverage.html
	@echo "####### HTML coverage report generated: coverage.html"


.PHONY: mocks-gen
mocks-gen:
	@echo "####### Generating mocks for project..."
	@mockery --all --recursive --dir ./internal --output ./internal/mocks --keeptree --case snake || { echo "####### ERROR: Mocks generation failed"; exit 1; }
	@echo "####### Mocks generated!"


#################################################################
# Other
#################################################################

.PHONY: clean
clean:
	@echo "####### Cleaning up temporary files..."
	@rm -f coverage.out coverage.html coverage_detailed.txt spotify-cli.exe
	@echo "####### All files were removed!"


#################################################################
# Developer sections
#################################################################

.PHONY: pre-commit
pre-commit: mocks-gen lint-fix build test test-coverage-detailed clean
	@echo "####### Validations succeeded, you can commit now!"
