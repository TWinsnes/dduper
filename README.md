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

## Release

This repo uses GoReleaser to release new version of the utility. 

Make sure `GITHUB_TOKEN` environment variable is set with a token that has repo access

Add a new tag following semver to the github repository and make sure it is pushed to GitHub

Create a new release with the following commend from the root of the repository

```sh
goreleaser release
```

More details on GoReleaser can be found in the [GoReleaser doco](https://github.com/goreleaser/goreleaser).