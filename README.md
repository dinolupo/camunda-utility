# Camunda Command Line Utility

This software aims to build a command line utility helper for the Camunda platform.

Build Executable
-----------
> MacOs/Linux
```bash
go build -o camunda-utility
```

> Cross Compiling for Windows
```bash
GOOS=windows GOARCH=amd64 go build -o bin/windows/camunda-utility.exe
GOOS=linux GOARCH=amd64 go build -o bin/linux/camunda-utility
GOOS=linux GOARCH=arm go build -o bin/linux-arm/camunda-utility
GOOS=darwin GOARCH=arm64 go build -o bin/darwin-arm/camunda-utility
```

Usage
-----------

> run a local camunda on port 8081 using the Docker at:

https://github.com/dinolupo/spring-camunda-template

> delete all definitions and instances
```bash
./camunda-utility deleteDefinition --host localhost --port 8081 --key "@all"
```

> help
```bash
./camunda-utility --help
Camunda Utility is a command line tool that permits to execute
        administrative tasks like deleting all definitions and instances.

Usage:
  camunda-utility [command]

Available Commands:
  completion       generate the autocompletion script for the specified shell
  deleteDefinition Delete Camunda definition and instances for a single or all process definitions
  help             Help about any command

Flags:
      --config string   config file (default is $HOME/.camunda-utility.yaml)
  -h, --help            help for camunda-utility
      --host string     Camunda Host (default "localhost")
      --port int        Camunda Port (default 8080)

Use "camunda-utility [command] --help" for more information about a command.
```

Testing
-----------

Run on the fly
```bash
go run main.go --host localhost --port 8081 deleteInstances --key @all
```

Unit-tests:
```bash
go test -v -race ./...
```

Run linter:
```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.27.0 golangci-lint run -v
```

CONTRIBUTE
-----------
 * write code
 * run `go fmt ./...`
 * run all linters and tests (see above)
 * run all examples (see above)
 * create a PR describing the changes

LICENSE
-----------
MIT

AUTHOR
-----------
Dino Lupo <dino.lupo at gmail.com>

BUILD REFERENCE
------------

https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04