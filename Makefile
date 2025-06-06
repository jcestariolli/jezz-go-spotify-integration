APP_NAME=spotify-cli

ifeq ($(OS),Windows_NT)
    EXE_EXT := .exe
    RUN_CMD := .\$(APP_NAME)$(EXE_EXT)
else
    EXE_EXT :=
    RUN_CMD := ./$(APP_NAME)p$(EXE_EXT)
endif

.PHONY: install-deps
install-deps:
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest


.PHONY: build
build:
	go build -o ./$(APP_NAME)$(EXE_EXT) ./cmd/$(APP_NAME)

.PHONY: run
run: build
	$(RUN_CMD)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-with-fix
lint-with-fix:
	golangci-lint run --fix

