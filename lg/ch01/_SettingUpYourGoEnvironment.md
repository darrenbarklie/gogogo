# Chapter 1: Setting Up Your Go Environment

Refer to official [installation instructions](https://go.dev/doc/install)

Add to _.zshrc_ file:

- `export GOPATH=$HOME/go`
- `export PATH=$PATH:$GOPATH/bin`

## Common commands

- `which go`
- `go version`
- `go run main.go`
- `go build main.go`
- `go build -o hello_world_v1.0.0 hello.go
- `go get .`

## Installing tools/packages

Go doesn't rely on a central repository, instead using code repositories:

- `go install github.com/rakyll/hey@lastest`
- `hey https://craftapplied.com`
- Update as is intalled with `@latest`

## Authoring Go

Format code to the singularly correct writing style with `go fmt`. This enforced style limits "format wars" and empowers tooling through convension.

Standardise imports to be ordered alphabetically, with unused removed and undeclared imports guessed, with `goimports -l -w .`.

- `-l` outputs list of affected files
- `-w` writes files in situ
- `.` defines directory, with subdirectories

Always run `go fmt` or `goimports` before compiling your code.
