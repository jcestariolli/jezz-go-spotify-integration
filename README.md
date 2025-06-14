# 🎶 jezz-go-spotify-integration

A Go-based integration project for interacting with the Spotify API. This project provides configuration loading, pagination utilities, mock generation, linting, building, and testing workflows to streamline development.

📌 **_Important_**: this project only authenticates with client credentials and, considering this, it doesn't integrate with Spotify's endpoints that access user information.

---


## 📚 Study Project Notice

**This project is primarily for study and learning purposes.** It's a personal endeavor to deepen understanding of Go programming, API integration (specifically with Spotify's API), and various development workflows like testing, linting, and mock generation. While functional, it might not adhere to all best practices for production-ready applications. Feel free to explore, learn from it, and provide feedback!

---

## 🚀 Project Description

This project is designed to facilitate integration with Spotify's API using Go. It includes:

* **Configuration management** with validation (YAML/JSON support) ⚙️
* **Utilities** for handling pagination parameters 📄
* **Automated mock generation** for interfaces 🤖
* **Linting and code quality checks** ✅
* **Build and test automation** with coverage reporting 🧪
* **Sample API calls** to demonstrate features and verify functionality. ✨

---

## 📁 Folder Structure

```
.
├── cmd
│   └── spotify-cli     # Main application entry point 🚀
│       ├── config      # Configuration files (e.g., config.yml, spotify_client_credentials.yml) 📁
│       ├── samples     # Contains sample code to demonstrate API interactions 💡
│       └── main.go     # Main application file ▶️
│── internal
│   ├── auth            # Implementations for Spotify authentication flows 🔑
│   ├── config          # Configuration structs, loaders, and validation logic 📝
│   ├── model           # Domain models and types used across the app 🧩
│   ├── resource        # Implementations for Spotify API integration for various features 🎵
│   ├── service         # Implementations of the business logic that will be executed before using resources 💼
│   ├── utils           # Utility functions (e.g., pagination validation) 🛠️
│   └── mocks           # Auto-generated mocks for testing 🤖
│── test
│   └── data            # Sample config files and test data 📊
├── .github
│   │── actions         # Configurations for common actions used on workflows ⚙️
│   └── workflows       # Configurations for github pipelines / workflows ⚙️
├── .gitignore          # File that specifies intentionally untracked files to ignore 🚫
├── .golangci.yml       # File with configuration for golangci-lint 📏
└── Makefile            # File that automates common development tasks 🛠️
```
This project's structure is organized to separate concerns. The primary application logic resides in `cmd/spotify-cli`, and internal libraries, models, utilities, authentication, and feature-specific implementations are within the `internal` directory. Root-level configuration files manage overall project settings and version control.

---

## ⚙️ Configuration

### Credentials and App Config

:information\_source: Configuration files must be placed inside the `cmd/spotify-cli/config` directory.

* The **`config.yml`** file contains Spotify URLs and other necessary settings. This file comes pre-filled and generally **should not be modified** unless absolutely necessary. ⚠️
* The developer **must create** a file named **`spotify_client_credentials.yml`** in the same folder. This file should contain your Spotify app `ID` and `secret` required to connect to the Spotify API. 🤫
* A sample credentials file named `spotify_client_credentials.yml.sample` is provided inside the `cmd/spotify-cli/config` folder. Developers should copy this sample file, rename it to `spotify_client_credentials.yml`, and fill in their own Spotify app credentials. ✍️

More details about configuring Spotify app credentials can be found on Spotify's documentation [Getting started with Web API](http://googleusercontent.com/spotify.com/4). 🔗

:warning: **Important**: **Do not commit your real credentials to version control.** Ensure `spotify_client_credentials.yml` is ignored by Git (add it to your `.gitignore` file). Make sure both files are properly configured to avoid connection or validation errors when running the application. 🔒

---

## 🛠️ Using the Makefile

The `Makefile` automates common development tasks, simplifying the workflow for developers:

* `make install-deps`
  Installs all Go dependencies required by the project. 📦
* `make tidy`
  Runs `go mod tidy` to clean up unused dependencies. ✨
* `make build`
  Compiles the project executable. 🏗️
* `make run`
  Compiles and then executes the project. 🏃
* `make lint`
  Runs `golangci-lint` to check code quality and style. 🔍
* `make lint-fix`
  Runs `golangci-lint --fix` to check code quality and style, and automatically apply suitable fixes. 🩹
* `make test`
  Executes all project tests. 🧪
* `make test-coverage`
  Runs tests with coverage reporting (excluding `model` and `mocks` packages). 📊
* `make test-coverage-detailed`
  Runs tests with detailed coverage reporting (excluding `model` and `mocks` packages), providing line-by-line coverage information. 📈
* `make mocks-gen`
  Generates mocks for interfaces used in tests. 🤖
* `make pre-commit`
  Runs a sequence of tasks: mocks generation, linting, build, and tests. Use this command before committing code to ensure quality and consistency. 💪

---

## ▶️ Getting Started

Follow these steps to set up and start using the project:

1.  **📝 Set up the configuration files** within the `cmd/spotify-cli/config` directory, ensuring all required fields are filled.
2.  **🏗️ Build the project**:
    ```bash
    make build
    ```
3.  **🧪 Run tests**:
    ```bash
    make test
    ```
4.  **📊 Run tests with coverage**:
    ```bash
    make test-coverage
    ```
5.  **✅ Run all pre-commit checks** after developing and before committing to ensure code quality:
    ```bash
    make pre-commit
    ```

---

## 📌 Notes

* Ensure you have all necessary dependencies installed, including `golangci-lint` and `mockgen`, and that you are using **Go 1.16+**. ✅
* Configuration validation leverages `go-playground/validator` to enforce required fields and data formats. 🔒
* The project utilizes Go generics for flexible configuration loading and validation. 💡