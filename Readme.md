# address-parser

Golang HTTP service wrapping [libpostal](https://github.com/openvenues/libpostal) to provide address segment detection. Libpostal is not bundled, it needs to be installed prior to building the project. Please consult libpostal's documentation for setup instructions. This project's GH Action workflow and Dockerfile might also provide some insights.

## Prepare

Tested with Go v1.17.

```bash
go mod tidy
```

## Build

```bash
go generate ./...
# generates *.gen.go files in ./api
go build
```

## Test

```bash
go test ./...
```

## Run

```bash
./address-parser
```
