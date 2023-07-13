---
title: Go Programming
author: Piyush Patel
---

# Overview of Go

- developed at Google and supported by Google
- Version 1 released in 2012
- Very efficient and high speed, compiles to machine code rather than JVM
- Built-in support for concurrency
- Syntax easy and concise
- Module system to expand functionality, built-in Production ready HTTP server
- Statically typed language but feels like dynamic
- Garbage collection and has fast compilation
- Built in tooling using `go` for built, run, test, install, etc.

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

# Prerequisites

- Basic familiarity with any other programming language
- VS Code or any other IDE setup
- Go SDK as well as Git (for pulling modules)

# Code

```java
public class Main {
    public static void main(String[] args) {
        System.println("Hello World")
    }
}
```

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```
