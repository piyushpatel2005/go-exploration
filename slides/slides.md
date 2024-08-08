---
title: Go Programming
author: 
    - Piyush Patel
---

## Overview of Go

::: incremental

- Developed at Google and supported by Google, version 1 in 2012
- Very efficient and high speed, compiles to machine code rather than JVM
- Built-in support for concurrency
- Syntax easy and concise
- Module system to expand functionality, built-in Production ready HTTP server
- Statically typed language but feels like dynamic
- Garbage collection and has fast compilation
- Built in tooling using `go` for build, run, test etc.
- Another bullet
- Next bullet

:::

::: wrapper

## Temporary slide

- A slide
- Second slide

## Nested Slide

- Nested slide content

:::

## Who uses Go?

- Google
- Facebook
- Netflix
- PayPal
- Many more...

## What is it used for?

- Kubernetes
- Docker
- Terraform

## Prerequisites

- Basic familiarity with any other programming language
- VS Code or any other IDE setup
- Go SDK as well as Git (for pulling modules)

## Code

```java {.numberLines .fragment}
public class Main {
    public static void main(String[] args) {
        System.println("Hello World")
    }
}
```

```go {.jsx .fragment}
package main
import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

## Installing Golang


### First Go Program

```go {.fragment data-line-numbers="|1|3|5,7|6" data-id="code-animation"}
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

```shell{.generic-code .fragment}
$ go run hello.go
```

```shell{.fragment}
Hello World
```