# dduper

utility to scan you local drive and delete any duplicate files it finds

It scans your drive alphabetically and will only keep the __1st__ copy, regardless of depth. e.g `a/b/c/d/e/f/g/h/xfile.txt` will be considered before `bfile.txt`

## Usage

Execute it from the location you would like to dedupe

```sh
./dduper
```

## Install

Either use `go install` or download the latest binary from Releases

```sh
go install github.com/twinsnes/dduper
```

## Testing

Tests are executed using go test, a make file wraps the command. See command below

```sh
make test
```

## Build

To build a local copy use the make wrapper for the golang build command

```sh
make build
```