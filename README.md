# Camunda Command Line Utility

This software aims to build a command line utility helper for the Camunda platform.

Usage
-----------

> run a local example camunda on port 8081 using the Docker folder:
```sh
docker compose -p camunda -f camunda-mariadb.yml up
```

> delete all instances of all definitions
```sh
./camunda-utility --host localhost --port 8081 deleteInstances --key @all
```

> delete all definitions and instances
```bash
./camunda-utility deleteDefinition --host localhost --port 8081 --key "@all"
```

> help
```bash
./camunda-utility --help
```

> delete Camunda stopped containers and volumes
```sh
docker rm -f $(docker ps -a -q)
docker volume rm $(docker volume ls -q)
```

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