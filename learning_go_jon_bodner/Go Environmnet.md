# The Go Workspace
Since the introduction of Go in 2009 there have been several chagnes in how Go developer's organize their code and dependencies. For modern Go development the rule is simple: you are free to organize your projects as you see fit.

It is still expected that there be a single workspace for third-party Go tools installed via `go install`. By default, this workspace is located in `$HOME/go`, you can use this default or specify a different workspace by setting the `$GOPATH` environment variable.

Whether or not you use the default location, it's a good practice to explicitly define `GOPATH` and to put the `$GOPATH/bin` directory in your executable path. 

# The go Command
Go ships with many development tools which can be accessed via the `go` command. They include a compiler, code formatter , linter, dependency manager, test runner, and more.

## go run and go build
`go run` and `go build` each take either a single Go file, a list of Go files, or the name of a package. When you enter `go run <file_name>` you'll notice that no binary has been saved there. The `go run` command does compile your code into a binary, however, the binary is built in a temporary directory. It then executes the binary from the temporary directory, and then deletes the binary after the program finishes. This makes `go run` useful for testing out small programs or using Go like a scripting language. 

Most of the time you want to build a binary for later use. That's where you use the `go build` command. `go build <file_name>` will create an executable file in the current directory with a name that matches the name of the file or package that was passed in. To change the name use the `-o` flag: `go build -o hell_world hello.go`

## Getting Third-Party Go Tools
Some people choose to distribute their Go programs as pre-compiled binaries, but tools written in Go can also be built from source and installed into your Go workspace via the `go install` command.

Go's method for publishing code is different than most other languages. Go developers don't rely on a central hosted service, like Maven Central or NPM. Instead, they share projects via their source code repositories. The `go install` command takes the location of the source code repository as an argument followed by an `@` and the version of the tool you want, use `@latest` for the latest version. It then download, compiles, and installs the tool in your $GOPATH/bin directory.

```shell
$ go install github.com/rakyll/hey@latest
go: downloading github.com/rakyll/hey v0.1.4
go: downloading golang.org/x/net v0.0.0-20181017193950-04a2e542c03f
go: downloading golang.org/x/text v0.3.0
```

The above example downloads `hey` and all of its dependencies, builds the program and installs the binary in your $GOPATH/bin directory. `hey` is a great Go tool that load tests HTTP servers. 

## Formatting Your Code
One of the chief design goals for Go was to create a language that allowed you to write code efficiently. This meant having a simple syntax and a fast compiler. Go enforces a standard format that makes it a great deal easier to write tools to manipulate source code, simplifies the compiler, and allows the creation of some clever tools for generating code.

The Go development tools include a command `go fmt` which automatically reformats your code to match the standard format. Always run `go fmt` before compiling your code.

## Linting and Vetting
While `go fmt` ensures your code is formatted correctly, it's just the first step in ensuring that your code is idiomatic and of high quality. All developers should read the [Code Review Comments page](https://github.com/golang/go/wiki/CodeReviewComments) on Go's wiki to understand what idiomatic Go code should look like.

To help enforce this style you can use a too like `golint`. Some of the changes it will suggest include properly naming variables, formatting error messages, and placing comments on public methods and types. 

```shell
$ go install golang.org/x/lint/golint@latest
$ golint ./... # runs golint over your entire project
```

Another class of errors that developers run into is code that is syntactically valid, bu there are mistakes that are not what you meant to do. This includes things like passing the wrong number of parameters to formatting methods or assigning values to variables taht are never used. The `go` tool includes a command called `go vet` to detect these kinds of errors.

# Makefiles
Modern software relies on repeatable, automatable builds that can be run by anyone, anywhere, at any time. The way to do this is to use some kind of script to specify your build steps. Here's a sample Makefile:

```makefile
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build hello.go
.PHONY:build
```

Each possible operation is called a *target*. The `.DEFAULT_GOAL` defines which target is run when no target is specified. Next are the target definitions. The word before the colon is the name of the target, any words after the target are other targets that must be run before the specified target runs. The tasks performed by the target are on the indented lines after the target. The `.PHONY` line keeps make from getting confused if you ever create a directory in your project with the same name as the target.

# Staying Up to Date
There are periodic update to the Go development tools. Go programs are native binaries that don't rely on a separate runtime, so you don't need to worry that updating your development  environment could cause your currently deployed programs to fail. You can have programs compiled with different version of Go running simultaneously on the same computer or VM.

# Next Section
[[Primitive Types and Declarations]]