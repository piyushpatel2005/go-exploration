# Go Modules and Packages

## Packages

Packages in Go are used to better organize your code. A package is a collection of Go files in a single directory. Every `.go` file in a package should start with a `package` declaration. The package declaration should be the first line in the file. The package declaration should be followed by the package name. 

```go
package <package_name>
```

This is what defines the scope of the package. The code defined in the same package can access and use types, constant, functions defined in the same package. This is true even if those are defined in different files in the same package. Let's verify this.

Create a go file named `packagedemo/calculator.go` with the following content.

```go
package main

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}
```

Here, I've defined two functions `Add` and `Subtract` which does simple math on two numbers. Notice, that the package name is `main`. 

Now, create another file named `packagedemo/main.go` with the following content.

```go
package main

import "fmt"

func main() {
	fmt.Println("Addition:", Add(10, 20))
	fmt.Println("Subtraction:", Subtract(20, 10))
}
```

Here, I'm using the methods defined in `calculator.go` file without any explicit import at the top. This is possible because the package of `main.go` file is same as `calculator.go` file. 

Now, to execute the code, run the following command.

```bash
cd packagedemo
go run *.go
```

### `main` package

In Go, the `main` package is a special package. It is the entry point of the program. When you run a Go program, the `main` package is executed first. Every `main` package should have a `main()` function which is the entry point of the program. If it doesn't exist, you will get an error. Conventionally, the `main()` method is defined in a file called `main.go` although it is not mandatory. 

Another important point is that a single directory cannot have multiple packages in it. For example, if you rename the package for `calculator.go`, you will see an error when you try to run the program. This is because at the moment `calculator.go` and `main.go` are in the same directory and they should have the same package name. Also, the name of the package should be the same as the directory name. This is an exception for the `main` package which can exist in any directory.

### Importing packages

To use code from another package, you need to import that package. You can import a package using the `import` keyword followed by the package path. The package path is the import path for the package. The package path is the full directory path of the package. You've already seen how to import packages like `strings`, `fmt` in previous tutorials. The full list of standard packages available can be found [here](https://pkg.go.dev/std). To import package like `os/user`, you have to specify full path like `import "os/user"`. 

```go
package main

import (
    "fmt"
	"os/user"
)

func main() {
    u, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Println(u.Username)
}
```

## Modules
Modules in Go are a collection of Go source files with a `go.mod` file in the root directory. This is higher level abstration to structure the code.  This promotes reusability of code. The `go.mod` file defines the module's path, which is the import path for the module. The `go.mod` file also lists the module's dependencies if you're using any third-party packages in your code.

### Creating a Module
You can create a module by using below command.

```shell
go mod init <module_name>
```

Let's create our first module with the directory `mymodule`.
    
```shell
mkdir mymodule
cd mymodule
go mod init mymodule
```

If you look at the contents of the `go.mod` file, you will see something like this.

```shell
module mymodule

go 1.21
```

This basically specifies the module path for this module and the Go version used. The module path is the import path for the module. This is used to resolve imports in your code.

In Go, it's a convention to base module path on a URL that is under your control. This is because the module path is used to resolve imports. If you're planning to publish your module, you should use a URL that you control. If you're not planning to publish your module, you can use any name you want. For example, the module path can be something like `github.com/<username>/<module_name>` or canonical url like `<module_name>.example.com`. You could use the command below to initialize a module with a specific path.

```shell
cd <module_name>
go mod init github.com/<username>/<module_name>
```

If you had defined the module using above command the `go.mod` file would look like this.

```shell
module github.com/<username>/<module_name>

go 1.21
```

Let's write some code in our module. Create `hello.go` file with the following content.

```go
package mymodule

func Hello(name string) string {
	return "Hello, " + name + "!"
}
```

Create another file under package `calculator` as `calculator/calculator.go` file with below content.

```go
package calculator

// Add function adds two integers and returns the result
func Add(a, b int) int {
	return a + b
}

// Subtract function subtracts two integers and returns the result
func Subtract(a, b int) int {
	return a - b
}

// Multiply function multiplies two integers and returns the result
func Multiply(a, b int) int {
	return a * b
}

// Divide function divides two integers and returns the result
func Divide(a, b int) int {
	return a / b
}
```

Now, create yet another file `types/Person.go` which defines few types we will use.

```go
package types

// Person struct represents a person with a name and age
type Address struct {
	Street string
	City string
	State string
	Country string
	Zip string
}

type Person struct {
	Name string
	Age  int
	Address Address
}
```

If you look at current directory structure, it should look like this.

```plaintext
tree .
.
├── calculator
│   └── calculator.go
├── go.mod
├── hello.go
└── types
    └── Person.go
```

### Adding External Dependencies

Let's add external dependency in our new module. Let's say we want to define the `Person` with unique ID. We can use `github.com/google/uuid` package to generate unique IDs. To add this dependency, run below command.

```shell
go get github.com/google/uuid
```

This modifies our `go.mod` file to include the dependency.

```go
module mymodule

go 1.21.5

require github.com/google/uuid v1.6.0 // indirect
```

If you see comment `// indirect` next to the dependency, it means that this is not directly used in your code but used by some other package that you are using.

To use this new package, you will modify your `types/Person.go` file to use the `uuid` package. This is external third-party dependency and not part of standard Go library, so you need to specify full path in import statement.

```go
import (
	"github.com/google/uuid"
)
```

Next, you can use the `UUID` type from the `uuid` package to generate unique IDs for the `Person` struct.

```go
type Person struct {
	Id      uuid.UUID
	Name    string
	Age     int
	Address Address
}
```

### Publishing Modules

You can modify the `go.mod` file with module name and version. This is useful when you want to publish your module. 

```go
module github.com/<username>/<module_name>

go 1.21.5

require (
    github.com/google/uuid v1.6.0
)
```

Next to publish your module, you can use below command.

```shell
go mod tidy
go test ./...
git add .
git commit -m "Initial commit"
git tag v0.1.0
git push origin v0.1.0
```

### Using This Module

To use this module in another project, you can use the `go get` command to download the module.

```shell
go get github.com/<username>/<module_name>
```

This will download the module to your `GOPATH` and you can use it in your code.

```go
package main

import (
    "fmt"
    "github.com/<username>/<module_name>/calculator"
    "github.com/<username>/<module_name>/types"
)

func main() {
    fmt.Println(calculator.Add(10, 20))
    fmt.Println(calculator.Subtract(20, 10))

    p := types.Person{
        Name: "John Doe",
        Age:  30,
        Address: types.Address{
            Street:  "123 Main St",
            City:    "New York",
            State:   "NY",
            Country: "US",
            Zip:     "12345",
        },
    }
    fmt.Println(p)
}
```