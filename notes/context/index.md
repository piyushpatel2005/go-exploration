# The Context Package

The context package is used to handle metadata for a goroutine. It is used to pass values between goroutines as well as to cancel a goroutine. The metadata in server's context can be the data required to process the request and the data about when to stop processing the request. For example, if the server is making a third party API call, you could set a timer and cancel the call if the timer expires.

## Overview

The context is an instance that implements `context.Context` interface. It's just an additional parameter to your function. The syntax for passing context into a function looks like below.

```go
func some_function(ctx context.Context, other_params) {
    // do something
}
```

## Creating Context

The `context` package has few functions to create context.

- `context.Background()`: This returns a variable of type `context.Context` with no values.
- `context.WithValue()`: This returns a variable of type `context.Context` with a value of type `value`.
- `context.TODO()`: This returns a variable of type `context.Context` with no values. This can be used as a placeholder for developer's convenience when they aren't sure where the context will come from.

An empty context can act as a starting point for a chain of contexts. Every time you want to pass a value into a context, you should use `context.WithValue()`.
