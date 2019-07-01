# Go Programming

- Go language is very fast. Machine language are CPU instructions and are very simple and straight forward and it runs directly on processors. Assembly languages are similar to machine language but there are English mnemonics. Go is a high level language. In compiled code or interpreted code, there is compile or interpretation step that needs to happen. Go language also has garbage collection for automatic memory management.
- It has simpler objects. Go is weakly object-oriented language. In Go, we don't have class, but we have `struct` with associated methods. In Go, we don't have constructor, inheritance or generics.
- It has efficient concurrency built into language. Concurrency is the management of multiple tasks at the same time. This requires management of task execution, so we need communication between tasks. `Goroutines` represent concurrent tasks. `Channels` are used to communicate between tasks. `Select` enables task synchronization.

## Installing Go

Download go tarball. `go1.12.6.linux-amd64.targ.gz`.

```shell
sudo tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```

The go code is organized into:

src - for source code files
pkg - containing packages (libraries)
bin - contains executables

Three must be one package called `main`. The codes can be grouped into different packages and can be imported in another source code file. When we build the main package, it creates executable program. Main package needs to have a `main()` function.

```go
package main
import "fmt" // importing a package
func main() {
  fmt.printf("Hello, world\n")
}
```

When we import a package, GoTool has to search directories specified by GOROOT and GOPATH environment variables. If we want to install some package in another directory, we have to change these two environment variables to find those extra packages. When we download go language, we get GoTool with it.

- `go build` - compiles the program. We can also provide bunch of packages as arguments. It creates executable for the main package. It will have `.exe` suffix for executable in Windows.
- `go doc` - prints documentation for a package.
- `go fmt` - format source code files
- `go get` - downloads packages and installs them
- `go list` - lists installed packages
- `go run` - compiles go files and runs the executable
- `go test` - runs tests using files ending in "_test.go"

## Variables

Names must start with letter. They can have number, underscore. Every variables need name and type. So, they need declaration.

```go
var x int # variable name type
var x,y int # declare two variables on same line
// Defining alias (alternate name for a type)
type Celsius float64
type IDnum int
var temp Celsius
var pid IDnum
var x int = 100 // Initialization
var x = 100 // Don't need to specify type of a variable
// Initialize after declaration
var x int // declare
x = 100 // initialize
var x int // default value is x =0
var x string // default value of string is x = ""
// Declaration and initilization in the same statement
x := 100 // type is inferred. This works only inside a function
```