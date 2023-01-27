# Introduction

This is the backend API behind [placeholder.app](https://placeholder.app). See the frontend repository [here](../frontend)

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
go run main.go
```

## Documentation

See [/documentation](./documentation/index.html)

Raw OpenAPI (Swagger) documentation included in [./.vscode/openapi.yml](.vscode/openapi.yml)

```bash
openapi-generator-cli generate --skip-validate-spec -i ./.vscode/openapi.yml -g html2 -o ./documentation
```

ThunderClient requests included in [./.vscode/thunder-tests](./.vscode/thunder-tests)

# Requirements & Dependencies

- Go 1.19
- [gin-gonic/gin](github.com/gin-gonic/gin)

# Credits

Loosely built from [ninjawerk/go-rest-boilerplate](https://github.com/ninjawerk/go-rest-boilerplate/).

- [ninjawerk](https://github.com/ninjawerk)
- [amattu2](https://github.com/amattu2)
- [All Contributors](../../contributors)
