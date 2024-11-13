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

### `context.TODO()`
This returns a variable of type `context.Context` with no values. This can be used as a placeholder for developer's convenience when they aren't sure where the context will come from.

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()
	printDetails(ctx)
}

func printDetails(ctx context.Context) {
	// do something
	fmt.Println("Context: ", ctx)
}
```

In above code, I have created a function `printDetails()` which takes `context.Context` object as a parameter. I have created a context object using `context.TODO()` which acts as a placeholder context and passed it to the function. 

### `context.Background()`
This returns a variable of type `context.Context` with no values. This is used when you plan to create a context from scratch.

```go
package main

import (
    "context"
    "fmt"
)

func main() {
    ctx := context.Background()
    printDetails(ctx)
}

func printDetails(ctx context.Context) {
    // do something
    fmt.Println("Context: ", ctx)
}
```

### `context.WithValue()`

This returns a variable of type `context.Context` with a value of type `value`. This way you can add data into the context. You can also use the data in the context by using `context.Value()` method. This is useful when you want to pass some data between functions. For example, the first authentication function can set the user ID in the context and the second function can use that user ID to fetch the user details. The function accepts three parameters: the parent context, the key, and the value. When you add new values into the context, it returns a new context with the new value but also keeps the old values. The key and the value can be of any type.

```go
package main

import (
    "context"
    "fmt"
)

func main() {
    ctx := context.WithValue(context.Background(), "user_id", 123)
    printDetails(ctx)
}

func printDetails(ctx context.Context) {
    // do something
    fmt.Println("Context: ", ctx)
    fmt.Println("User ID: ", ctx.Value("user_id"))
}
```

In this code, I have created a context object using `context.Background()` and added a value `user_id` with value `123` using `context.WithValue()`. I have passed this context object to the `printDetails()` function. In the `printDetails()` function, I have printed the context object and the value of `user_id` using `ctx.Value()` method.

### `context.Done()`

The `context.Context` also inclues a method to verify whether the context has expired or not. The `context.Done()` method returns a channel which is closed when the context is cancelled or times out. In this case, a `select` statement can be used to check if the context is cancelled or not.

### `context.withCancel()`
This returns a new context and a cancel function. The cancel function can be used to cancel the goroutine. This can be useful when you're making a third party API call and you want to cancel the call if it takes too long. You do not want to wait for the API call to complete if it takes too long. The context also has a method `context.WithTimeout()` which returns a new context and a channel. The channel will be closed after the timeout and the goroutine can be cancelled when the channel is closed. This helps with timed out requests.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go process(ctx)
	time.Sleep(1 * time.Second)
	cancel()
}

func process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context is cancelled")
			return
		default:
			fmt.Println("Processing data")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
```

In above code, I have created a context object using `context.WithCancel()` and passed it to the `process()` function. In the `process()` function, I have used a `select` statement to check if the context is cancelled or not. If the context is cancelled, it will print "Context is cancelled" and return from the function.

### `context.WithTimeout()`
This returns a new context and a channel. The channel will be closed after the timeout and the goroutine can be cancelled when the channel is closed. This helps with timed out requests. This is similar to `context.WithCancel()` but with a timeout. In the above code, you could replace `context.WithCancel()` with `context.WithTimeout()` and pass the timeout duration as `1*time.Second`.

```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
```

### `context.WithDeadline()`

This returns a new context and a channel. The channel will be closed after the deadline time has passed and the goroutine can be cancelled when the channel is closed. This helps with timed out requests.

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
    go process(ctx)
    time.Sleep(2 * time.Second)
    cancel()
}

func process(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Context is cancelled")
            return
        default:
            fmt.Println("Processing data")
            time.Sleep(100 * time.Millisecond)
        }
    }
}
```

In this case, I have created a context object using `context.WithDeadline()` and passed the deadline time as `time.Now().Add(1*time.Second)`. I have passed this context object to the `process()` function. In the `process()` function, I have used a `select` statement to check if the context is cancelled or not. If the context is cancelled, it will print "Context is cancelled" and return from the function.

> Although context provides a way to store data, it is not recommended to store everything in the context. It is recommended to store only request-scoped data in the context which is used across multiple functions. For example, the user ID, request ID, etc. are request-scoped data and can be stored in the context. The context should not be used to store application-wide data.