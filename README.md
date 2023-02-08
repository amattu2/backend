# Introduction

This is the backend API behind [placeholder.app](https://placeholder.app). See the frontend repository [here](../../../frontend)

[![Go Tests](https://github.com/placeholder-app/backend/actions/workflows/go-tests.yml/badge.svg)](https://github.com/placeholder-app/backend/actions/workflows/go-tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/placeholder-app/backend)](https://goreportcard.com/report/github.com/placeholder-app/backend)
[![Dependency Review](https://github.com/placeholder-app/backend/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/placeholder-app/backend/actions/workflows/dependency-review.yml)

## Examples

Sizes supported are 30px thru 4000px. These are just example sizes with default options.

<details>
  <summary>Image Examples</summary>

  ![img](https://api.placeholder.app/image/30x30)

  ![img](https://api.placeholder.app/image/70x70)

  ![img](https://api.placeholder.app/image/120x120)

  ![img](https://api.placeholder.app/image/160x160)

  ![img](https://api.placeholder.app/image/200x200)

  ![img](https://api.placeholder.app/image/245x245)

  ![img](https://api.placeholder.app/image/310x310)

  ![img](https://api.placeholder.app/image/450x450)

  ![img](https://api.placeholder.app/image/650x650)

</details>

# Usage

## Setup

```bash
git clone https://github.com/placeholder-app/backend
```

```bash
cd backend/
```

```bash
go install
```

## Building

To build all binaries from the source

```bash
make build
```

To build a specific binary

```bash
make build_[platform]
```

Platforms: `linux`, `mac`, `freebsd`, `windows`

All of the above commands output to the `/bin` folder by default. To run the application, simply execute the build

```bash
./bin/[platform]
```

To clean the workspace

```bash
make clean
```

## Running Directly

If you choose to run without compiling, you can refer to this section.

Run with default options (For local testing)

```bash
go run main.go
```

Run with configuration (See [main.go](main.go) for options)

```bash
PORT=80 go run main.go
```

Run with SSL

```bash
SSLCERT=./cert.pem SSLKEY=./key.pem PORT=443 go run main.go
```

**Note:** On Linux, you may need to run go as `sudo` for ports 80 or 443

# Testing & Documentation

## Testing

The included makefile provides options to run golang test cases, and to build the openapi documentation.

To run the included Golang unit tests, with coverage,

```bash
make tests
```

**Note:** This will not output any files. See the commandline output

## Building OpenAPI Docs

The OpenAPI (swagger) file is in [.vscode/openapi.yml](.vscode/openapi.yml).

To build the HTML documentation file (found on [docs.placeholder.app](https://docs.placeholder.app)), see below.

```bash
make docs
```

Alternatively, you can run the command directly

```bash
openapi-generator-cli generate --skip-validate-spec -i ./.vscode/openapi.yml -g html2 -o ./docs
```

Both of the above commands output the HTML documentation inside ./docs

## HTTP Request Simulation

If you use VisualStudio Code, the ThunderClient request repository is available in [.vscode/thunder-tests](.vscode/thunder-tests).

You can simulate requests to the API, or run HTTP-based testing.

# Requirements & Dependencies

- Go 1.19
- [gin-gonic/gin](github.com/gin-gonic/gin)
- openapi-generator-cli (For documentation only)

# Credits

Loosely built from [ninjawerk/go-rest-boilerplate](https://github.com/ninjawerk/go-rest-boilerplate/).

- [ninjawerk](https://github.com/ninjawerk)
- [amattu2](https://github.com/amattu2)
- [james-elicx](https://github.com/james-elicx)
- [All Contributors](../../contributors)
