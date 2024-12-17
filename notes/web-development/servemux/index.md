# ServeMux

ServeMux is a struct in the `net/http` package that implements the `http.Handler` interface. It is used to route incoming HTTP requests to their respective handler functions.

In Go `serveMux`, it will always match longer URL patterns first. They have higher precedence over the shorter ones. If there are multiple patterns matching the URL, the one with longest pattern will be executed. URL path are by default sanitized. That is if they have trailing slashes or double dots, they are automatically redirected to the correct path. So, if you have a URL like `/about/../contact`, it will be redirected to `/contact`. The client will receive a `301 Moved Permanently` status code redirecting them to the correct URL.

You can also match based on the host names in the URL pattern. This can be useful if you want to define different handlers for different subdomains. For example, you can have `api.example.com` and `www.example.com` and have different handlers for each of them.

```go
func api(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("API page"))
}

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Home page"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("api.example.com/", api)
    mux.HandleFunc("www.example.com/", home)
    port := ":8000"

    log.Print("Starting server on ", port)

    err := http.ListenAndServe(port, mux)
    log.Fatal(err)
}
```

During pattern matching, the host specific patterns are matched first and then the general patterns.

### Limitations of ServeMux

Go's serveMux does not support routing based on request methods. If you come from other frameworks like Express, Django, Rails, you might have come across this feature. In Go, your handler has to handle the different request methods itself. You can access the request method using `r.Method` and then write the logic based on the method. `serveMux` also does not support routing based on regular expressions.

## Default ServeMux

If you don't want to create a new `serveMux` instance, you can use the default `serveMux` provided by the `http` package. This is the default `serveMux` that is used by the `http.ListenAndServe` function.

For example, if you do not specify the `serveMux` instance in the `http.ListenAndServe` function, it will use the default `serveMux`. You can also access this default `serveMux` using the `http.DefaultServeMux` global variable. In this case, you can use the `http.HandleFunc` method to add handlers to the default `serveMux`.

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
    http.HandleFunc("/", home)
    http.HandleFunc("/about", about)
    http.HandleFunc("/contact", contact)
    port := ":8000"

    log.Print("Starting server on ", port)

    err := http.ListenAndServe(port, nil)
    log.Fatal(err)
}
```

Here, `http.DefaultServeMux` is used to add handlers to the default `serveMux`. As this is a global variable, it can be accessed from anywhere in your project and you can add handlers to it. This also can cause an issue if any third party package registers routes with the default `serveMux` which can be security risk. So, it's better to use a custom `serveMux` instance.

## Handling 404 Not Found

If you want to create a default 404 page which you want to display if none of the routes match, then you have to include this as a condition in the `home` function.

```go
func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Home page"))
}
```

This function will ensure that it's exactly `/` that will display the `Home page` and anything else will display a 404 page.

With these change, if you now visit `/something`, you will get `404 page not found` message.

