# Go `time` package

Go has a built-in package called `time` which provides functionality for working with date and time. This package provides several types and functions to work with date time. Below are important types defined in the `time` package.

- `time.Time`: This is used to represent a point in time.
- `time.Duration`: Used to represent duration of time.
- `time.Timer`: Used to create a timer.
- `time.Month`: Used to represent a month. It's basically an enumeration (1 to 12) defined in the package.
- `time.Weekday`: Used to represent a day of the week. This is also defined as an enumeration 0 to 6.
- `time.Location`: This is used to map time instants to the zone in which they occurred.

The `time.Time` type represents a point in time with upto nanoseconds accuracy. It has several methods to work with time. 

## Creating Time

You can create a time object using `time.Now()` function. This function returns the current time in UTC timezone. 

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println(t) // 2024-10-17 15:04:05.123456789 +0000 UTC m=+0.000000001
}
```

You could also create a time object using `time.Date()` function. This function takes year, month, day, hour, minute, second, nanosecond and location as arguments and returns a time object.

```go
t := time.Date(2024, time.October, 17, 15, 4, 5, 123456789, time.UTC)
fmt.Println(t) // 2024-10-17 15:04:05.123456789 +0000 UTC
```

## Time Methods

### Getting Year, Month, Day, Hour, Minute, Second

You can get year, month, day, hour, minute, second using their respective methods. These methods return an integer value or enum.
The `Zone` method returns the timezone such as (EST, CST) and offset (in seconds) of the time object.

```go
t := time.Date(2024, time.October, 17, 15, 4, 5, 123456789, time.UTC)
fmt.Println(t) // 2024-10-17 15:04:05.123456789 +0000 UTC

fmt.Println(t.Year()) // 2024
fmt.Println(t.Month()) // October
fmt.Println(t.Day()) // 17
fmt.Println(t.Hour()) // 15
fmt.Println(t.Minute()) // 4
fmt.Println(t.Second()) // 5
fmt.Println(t.Nanosecond()) // 123456789
fmt.Println(t.Weekday()) // Wednesday
fmt.Println(t.Zone()) // UTC 0
```

### Adding and Subtracting Time

You can add and subtract time using `Add` method with positive or negative arguments. These methods take a `time.Duration` object as argument and return `Time` object.

```go
t := time.Date(2024, time.October, 17, 15, 4, 5, 123456789, time.UTC)
fmt.Println(t) // 2024-10-17 15:04:05.123456789 +0000 UTC

t1 := t.Add(24 * time.Hour)
fmt.Println(t1) // 2024-10-18 15:04:05.123456789 +0000 UTC

t2 := t.Add(-24 * time.Hour)
fmt.Println(t2) // 2024-10-16 15:04:05.123456789 +0000 UTC
```

### Comparing Time

You can use methods `Before`, `After` and `Equal` to compare time objects. All these methods take a `Time` object as argument and return a boolean value.

```go
t := time.Date(2024, time.October, 17, 15, 4, 5, 123456789, time.UTC)
t1 := t.Add(24 * time.Hour)
t2 := t.Add(-24 * time.Hour)
fmt.Println(t1.Before(t2)) // false
fmt.Println(t1.After(t2))  // true
fmt.Println(t1.Equal(t2))  // false
```

### Formatting Time

You can format time using `Format` method. This method takes a string as argument which specifies how the time should be formatted. Below are some of the common formats. You already saw the default format when you printed the time object.

```go
t := time.Date(2024, time.October, 17, 15, 4, 5, 123456789, time.UTC)
fmt.Println(t.Format("2006-01-02 15:04:05"))    // 2024-10-17 15:04:05
fmt.Println(t.Format("2006-01-02 03:04:05 PM")) // 2024-10-17 03:04:05 PM
fmt.Println(t.Format("02/01/2006"))             // 17/10/2024
fmt.Println(t.Unix())                           // 1729177445
fmt.Println(t.UnixMilli())                      // 1729177445123
```

### Parsing Time

You can parse a string to time object using `time.Parse` function. This function takes a layout and a string as arguments and returns a `Time` object. If it fails to parse the string, it returns an error so it's a good idea to check the error.

```go
t, err := time.Parse("2006-01-02", "2024-10-17")
fmt.Println(t, err) // 2024-10-17 00:00:00 +0000 UTC <nil>
```

## Working with Location

You can get the location of the time object using `Location` method. This method returns a `Location` object. 

You can get new location using `LoadLocation` method which takes in a string representing the timezone location. 

You can also set the location of the time object using `In` method. This method takes a `Location` object as argument and returns a `Time` object.

```go
t := time.Date(2024, time.October, 17, 15, 4, 5, 123456789, time.UTC)
fmt.Println(t.Location()) // UTC

loc, _ := time.LoadLocation("America/New_York")
t1 := t.In(loc)
fmt.Println(t1) // 2024-10-17 11:04:05.123456789 -0400 EDT

fmt.Println(t.In(time.Local)) // 2024-10-17 11:04:05.123456789 -0400 EDT
fmt.Println(t.In(time.UTC))   // 2024-10-17 15:04:05.123456789 +0000 UTC
fmt.Println(t.In(time.LoadLocation("Asia/Kolkata"))) // 2024-10-17 20:34:05.123456789 +0530 IST
```

## Duration

A duration represents the elapsed time between two instants as an int64 nanosecond count. You can create a duration object using `time.Duration` type. This type has constants defined for common durations like `time.Second`, `time.Minute`, `time.Hour` etc.

You can also calculate the duration between two time objects using `-` operator. This will return a `Duration` object.

- `Sub` method is used to calculate the difference between two time objects. This method returns a `Duration` object.
- `Seconds` method is used to get the duration in seconds.
- `Minutes` method is used to get the duration in minutes.
- `Since` method is used to get the time lapsed since the time object to current time.
- `Abs` method is used to get the absolute duration between two time objects.
- `Round` method is used to round the time object to the nearest duration based on time duration.

```go
t1 := time.Date(2024, time.October, 17, 15, 4, 5, 123456789, time.UTC)
t2 := time.Date(2024, time.October, 17, 15, 5, 6, 123456789, time.UTC)
fmt.Println(t2.Sub(t1)) // 1m1s
fmt.Println(t2.Sub(t1).Seconds()) // 61
fmt.Println(t2.Sub(t1).Minutes()) // 1
fmt.Println(time.Since(t1)) // time lapsed since t1 to current time
fmt.Println(t2.Sub().Abs(t1)) // 1m1s
fmt.Println(t1.Round(time.Minute)) // 2024-10-17 15:04:00 +0000 UTC
```