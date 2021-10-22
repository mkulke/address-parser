# address-parser

Golang HTTP service wrapping [libpostal](https://github.com/openvenues/libpostal) to provide address segment detection. Libpostal is not bundled, it needs to be installed prior to building the project. Please consult libpostal's documentation for setup instructions. This project's GH Action workflow and Dockerfile might also provide some insights.

## Prepare

Tested with Go v1.17 on Linux and MacOS.

```bash
go mod tidy
```

If libpostal is installed to a non-default path `PKG_CONFIG_PATH` and `LD_LIBRARY_PATH` need to be adjusted accordingly:

```bash
# LIBPOSTAL_HOME should point to the installation folder
export PKG_CONFIG_PATH=$LIBPOSTAL_HOME/lib/pkgconfig
export LD_LIBRARY_PATH=$LIBPOSTAL_HOME/lib
```

## Build

### Local

```bash
go generate ./...
# generates *.gen.go files in ./api
go build
```

### Docker

```bash
docker build -t address-parser .
```

## Test

```bash
go test ./...
```

## Run

```bash
./address-parser
```
