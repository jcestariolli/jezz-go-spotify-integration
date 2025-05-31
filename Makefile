.PHONY: build run tidy test

APP_NAME=spotify-integration

# Detect OS
ifeq ($(OS),Windows_NT)
    EXE_EXT := .exe
    RUN_CMD := .\$(APP_NAME)$(EXE_EXT)
else
    EXE_EXT :=
    RUN_CMD := ./$(APP_NAME)p$(EXE_EXT)
endif

build:
	go build -o $(APP_NAME)$(EXE_EXT) ./app

run: build
	$(RUN_CMD)

tidy:
	go mod tidy

test:
	go test ./...
