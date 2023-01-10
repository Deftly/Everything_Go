# The Go Workspace
Since the introduction of Go in 2009 there have been several chagnes in how Go developer's organize their code and dependencies. For modern Go development the rule is simple: you are free to organize your projects as you see fit.

It is still expected that there be a single workspace for third-party Go tools installed via `go install`. By default, this workspace is located in `$HOME/go`, you can use this default or specify a different workspace by setting the `$GOPATH` environment variable.

Whether or not you use the default location, it's a good practice to explicitly define `GOPATH` and to put the `$GOPATH/bin` directory in your executable path. 

# The go Command
Go ships with many development tools which can be accessed via the `go` command. They include a compiler, code formatter, linter, dependency manager, test runner, and more.

## go run and go build
`go run` and `go build` each take either a single Go file, a list of Go files, or the name of a package. `go run` 