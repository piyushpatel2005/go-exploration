# Handling URL Query Strings

Query strings are a way to pass data to the server through the URL. They are appended to the end of the URL after a `?` character. The query string is a key-value pair separated by an `=` character. Multiple key-value pairs are separated by an `&` character.

```plaintext
http://test.com/path?name=John&age=30
```

In this case, the query string is `name=John&age=30`. The key-value pairs are `name=John` and `age=30`. The key is `name` and the value is `John`. The key is `age` and the value is `30`.

If you need to retrieve these query parameters in your route handler, you can use the `r.URL.Query().Get()` method. The `Get()` method takes the anem of the query param and returns the value of the query param as a string. If the query param is not found, it returns an empty string.

```go