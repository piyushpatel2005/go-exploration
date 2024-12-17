# Project Setup

In order to start working with a new project, you need to have your project organized into proper structure. This is usually done using a directory structure.

```shell
mkdir -p /code/<project_name>
cd /code/<project_name>
go mod init github.com/<username>/<project_name>
```

This creates a basic project structure with `go.mod` file inside this directory. This file is used to manage dependencies for your project and at this stage it should look like this.

```go
module github.com/<username>/<project_name>

go 1.21.5
```

While developing web applications, you need to handle different HTTP requests using different handler functions. These are also called views or controllers in other programming languages and/or frameworks. They are simply piece of code that executes some logic when a request is made to the web server and responds with HTTP response.

Now you might wonder if you have multiple paths then how does the web server know which function or handler to execute for each of those. This is where router comes into play. Routers provide mapping of URL paths to handler functions. In Go, routing is handled using `serveMux` which is a HTTP request multiplexer.

## Create Web server
Create `main.go` file in the root of the project.

```shell
touch main.go
```

Add the following code into the file.

```go
package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	port := ":8000"

	log.Print("Starting server on ", port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
```

Now, if you visit `http://localhost:8000` in your browser, you should see `Hello there` printed on the screen.
If you want to stop the server, you can press `Ctrl + C` in the terminal.

## Handling Multiple Routes

If you have multiple route handlers, you can add them as functions and map them using `mux.HandleFunc` method. One thing to note about URL patterns is that `servemux` supports two types of URL patterns.
1. Fixed paths: `/about` and `/contact` are fixed paths. They will only match if the URL path is exactly `/about` or `/contact`.
2. Subtree paths: `/about/` and `/contact/` are subtree paths. Notice that these paths have trailing slash (`/`). They will match if the URL path is `/about` or `/about/` or `/about/anything`.

```go
package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"))
}

func about(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("About page"))
}

func contact(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Contact page"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/about", about)
    mux.HandleFunc("/contact", contact)
    port := ":8000"

    log.Print("Starting server on ", port)

    err := http.ListenAndServe(port, mux)
    log.Fatal(err)
}
```

The root path `/` will match any URL path that is not matched by any other handler. You can try navigating to `http://localhost:8000/something` and you will still see the `Home page` displayed. 
