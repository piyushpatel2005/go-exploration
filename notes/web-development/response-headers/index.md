# Working with Headers

This tutorial will teach you how to work with response headers and how to set different headres in the HTTP response in Go.

If you're working with HTTP servers, you might have come across the need to set headers in the response. Headers are key-value pairs that are sent along with the response. They provide additional information about the response or the server. For example, you can set headers to specify the content type of the response, the server type, the cache control, you can do so using response headers.

## Setting Headers

In Go, you can set headers in the response using the `Header()` field of the `http.ResponseWriter` interface. This field is a map of string to slice of strings. You can set headers by adding key-value pairs to this map.

Below is an example of how you can set the `Content-Type` header in the response.

```go
func home(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte("<h1>Home page</h1>"))
}
```

This will render correctly in the browser as heading text. If the `Content-Type` header is not set, the browser will render the response as plain text.

You can also set the status code of the response using the `WriteHeader()` method of the `http.ResponseWriter` interface. This method takes an integer as an argument which is the status code of the response.

```go
func home(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
	//w.WriteHeader(200)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("<h1>Home page</h1>"))
}
```

Here, you can see that you can set the status code using `WriteHeader()` method by passing the exact status code or you can use the constants provided by the `http` package for various status codes. The second option is better because it makes the code more readable and less error-prone due to a typo.

> **Note:** It's only possible to call the `WriteHeader()` method once per response. If you try to call it again, Go will log a warning message. Also, if you don't set the status code explicitly, Go will set it to `200 OK` by default. If you change response headers after using `WriteHeader()`, it will have no effect on the response headers.

The status codes provide information abotu the status of the request. It's important to set the correct status code in the response to provide the correct information to the client. For example, if you're creating a resource and your method cannot handle other request methods, you should return `405 Method Not Allowed` status code for those methods as shown below.

```go
func create(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("Method not allowed"))
        return
    }
    // Create the resource
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Resource created"))
}
```

Again, this above snippet, I am using some of the constants provided by the `http` package such as `http.MethodPost` and `http.StatusMethodNotAllowed`.

### Default Headers

By default, Go generates some headers in the response. For example, it sets the `Date` header with the current date and time, the `Content-Length` header with the length of the response body, and the `Content-Type` header with the content type of the response. You can override these headers by setting them explicitly in the response. The `Content-Type` header is set to the correct value based on the response body. Basically Go will sniff the content of response body using `http.DetectContentType()` function and set the `Content-Type` header accordingly. If it can't guess the content type, it will set the `Content-Type` header to `application/octet-stream`. It sometimes fails to guess the content type correctly, so it's better to set the `Content-Type` header explicitly.

## Setting Multiple Headers

You can set multiple headers in the response by calling the `Set()` method multiple times on the `Header()` field of the `http.ResponseWriter` interface.

```go
func home(w http.ResponseWriter, r *http.Request) {
	// Set headers
    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("Cache-Control", "no-cache")
	
	// Add to existing header
	w.Header().Add("Cache-Control", "max-age=21500000")
	
	// Get all values for a header
	w.Header().Values("Cache-Control")
	
	// Delete a header key
	w.Header().Del("Cache-Control")
	
	// Get the value of a header key
	w.Header().Get("Content-Type")
    w.Write([]byte("<h1>Home page</h1>"))
}
```

If you need to delete the system generated headers such as `Date`, `Content-Length`, etc., you cannot do so using the `Del()` method. You can only delete the headers that you have set explicitly.

```go
w.Header()["Date"] = nil
```