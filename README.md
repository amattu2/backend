# Introduction

This is the backend API behind [placeholder.app](https://placeholder.app). See the frontend repository [here](../../../frontend)

# Usage

## Install

```bash
git clone https://github.com/placeholder-app/backend
```

```bash
cd backend/
```

```bash
go install
```

## Running

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

## Documentation

OpenAPI (Swagger) documentation included in [.vscode/openapi.yml](.vscode/openapi.yml)

```bash
openapi-generator-cli generate --skip-validate-spec -i ./.vscode/openapi.yml -g html2 -o ./documentation
```

ThunderClient requests included in [.vscode/thunder-tests](.vscode/thunder-tests)

# Requirements & Dependencies

- Go 1.19
- [gin-gonic/gin](github.com/gin-gonic/gin)

# Credits

Loosely built from [ninjawerk/go-rest-boilerplate](https://github.com/ninjawerk/go-rest-boilerplate/).

- [ninjawerk](https://github.com/ninjawerk)
- [amattu2](https://github.com/amattu2)
- [james-elicx](https://github.com/james-elicx)
- [All Contributors](../../contributors)
