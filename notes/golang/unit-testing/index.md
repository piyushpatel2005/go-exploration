# Unit Testing in Go

Unit testing is a software testing technique where individual units or components of a software are tested in isolation from the rest of the application. The purpose of unit testing is to validate that each unit of the software performs as designed. In Test driven development, the developer first writes test case and then writes the actual code that makes the test succeed. In Go, you can write unit tests for your code using the `testing` package.

There are different types of testing levels. The details for these are beyond the scope of this tutorial, but below is a list of different types of testing levels.

- **Unit Testing**: Testing individual units or components of a software.
- **Integration Testing**: Testing the integration of different components of a software.
- **System Testing**: Testing the entire system as a whole.
- **Acceptance Testing**: Testing the software to check if it meets the requirements.
- **Regression Testing**: Testing the software to check if new changes have affected the existing functionality.

## Writing Unit Tests in Go

Go provides a `testing` package which allows you to write unit tests for your code. First of all, for any testing code, you should use the convention of naming the test file with `_test.go` suffix. This is because the Go compiler ignores the files with `_test.go` suffix when building the application. For instance, if you have a file named `Person.go` and you want to write tests for it, you should create a file named `Person_test.go` in the same directory.

In order to write unit tests, let's first define a new module. Create a directory named `testing-demo` and navigate to it. Next, initialize a new module using the `go mod init` command.

```shell
mkdir testing-demo
cd testing-demo
go mod init testing-demo
```

Next, create a new file named `Greeter.go` with the following content.

```go
package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func MorningGreet(name string) string {
	return "Good morning, " + name + "! How are you doing today?"
}

func main() {
	fmt.Println(Hello("Chris"))
}
```

Now, create a new test file for this `Greeter.go` file. As mentioned earlier, the file name should have suffix `_test.go`, so create a new file named `Greeter_test.go` with the first test for `Hello()` function. Here, because both these files are the in the same package, we do not need explicit import at the top. You can directly call `Hello()` method from the `Greeter.go` file.

```go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Here, first of all, I have imported `testing` module which provides functionalities to write unit tests. Next, I have defined a function `TestHello()`. Each test function takes a pointer of type `testing.T`. This is a built-in type in the `testing` package which allows you to manage testing and record errors during testing. Next, you call the actual `Hello()` function by passing a name into it and store the returned value in a variable. If the returned value `got` do not match the expected value `want` then you print the error using testing module's `Errorf` function which accepts formatted string with placeholders and you can print any kind of error. This is usually a convention to print error messages like this.

## Executing Tests

Now, in order to run the tests, you can navigate to this directory and run the following command.

```shell
go test ./...
ok      testing-demo    0.251s
```

This will run the tests in the current directory as well as all subdirectories. However, it doesn't show what was executed. If you want to see the output of each test, you can run the following command.

```shell
go test -v ./...
```

If you want to execute only specific tests, you can run the following command.

```shell
go test -v -run TestHello
```

or if you want to run tests related to specific package, you can run the following command.

```shell
go test -v Greeter.go Greeter_test.go
```

In this case, you need to specify both the files.

## New Package with Multiple Tests

Let's write new package `Calculator.go` with four methods listed below.

```go
package main

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}
```

Now, to test these methods, create a new file named `Calculator_test.go` with the following content.

```go
package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestAdd2(t *testing.T) {
	got := Add(-1, 2)
	want := 1
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
```

Here, I have added two tests for `Add()` method. Both these tests look very similar expect the input and output values are different. If we had to write like 10 different tests for this function, it would be very repetitive and inefficient. 

### Avoiding Cached Results

By default Go will cache the results of previous test runs and only re-run tests that have changed. You can notice this if you try to skip one of the tests. In order to skip a test, you can use `t.Skip()` which will need string argument specifying reason for skipping. Here, I have modified `want` value to `2` in `TestAdd2()` function so this test should fail if run.

```go
func TestAdd2(t *testing.T) {
	got := Add(-1, 2)
	want := 2
	if got != want {
		t.Skip("Result is not as expected")
	}
}
```

Now, if you try rerunning the tests, you will notice that it still runs all the tests.

```shell
go test -v ./...
```

At the end of the output, you would see `(cached)`.

```plaintext
PASS
ok      testing-demo    (cached)
```

If you want to avoid cached results, you can use `-count` flag with `go test` command.

```shell
go test -v -count=1 ./...
```

## Table Driven Tests

In above example, we have written multiple tests for `Add()` method. This is a common pattern in Go to write multiple tests for a single method. However, writing multiple tests for a single method can be repetitive and inefficient. Go provides a way to write table-driven tests which allows you to write multiple tests for a single method in a more efficient way. Let's see how we can write table-driven tests for `Divide()` method.

```go
func TestDivide(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{a: 4, b: 2, expected: 2},
		{a: 2, b: 4, expected: 0},
		{a: 10, b: 2, expected: 5},
	}

	for _, tc := range tests {
		got := Divide(tc.a, tc.b)
		if got != tc.expected {
			t.Errorf("Divide(%d, %d): got %d, want %d", tc.a, tc.b, got, tc.expected)
		}
	}
}
```

In this case, I have defined a slice of anonymous struct which contains three fields `a`, `b`, and `expected`. I have defined three test cases for `Divide()` method. Next, I have used a `for` loop to iterate over each test case and call the `Divide()` method with `a` and `b` values and store the result in `got` variable. If the result `got` does not match the expected value `tt.expected`, then I print the error message using `t.Errorf()` function.
