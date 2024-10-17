package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Creating Time")
	t := time.Now()
	fmt.Println(t)
	t = time.Date(2024, time.October, 17, 15, 4, 5, 123456789, time.UTC)
	fmt.Println(t) // 2024-10-17 15:04:05.123456789 +0000 UTC

	fmt.Println("Getting time components")
	fmt.Println(t.Year())       // 2024
	fmt.Println(t.Month())      // October
	fmt.Println(t.Day())        // 17
	fmt.Println(t.Hour())       // 15
	fmt.Println(t.Minute())     // 4
	fmt.Println(t.Second())     // 5
	fmt.Println(t.Nanosecond()) // 123456789

	fmt.Println("Adding or subtracting time")
	t1 := t.Add(24 * time.Hour)
	fmt.Println("After 24 hours", t1) // 2024-10-18 15:04:05.123456789 +0000 UTC

	t2 := t.Add(-24 * time.Hour)
	fmt.Println("Before 24 hours", t2) // 2024-10-16 15:04:05.123456789 +0000 UTC

	fmt.Println("Comparing Time")
	fmt.Println(t1.Before(t2)) // false
	fmt.Println(t1.After(t2))  // true
	fmt.Println(t1.Equal(t2))  // false

	fmt.Println("Formatting time")
	fmt.Println(t.Format("2006-01-02 15:04:05"))    // 2024-10-17 15:04:05
	fmt.Println(t.Format("2006-01-02 03:04:05 PM")) // 2024-10-17 03:04:05 PM
	fmt.Println(t.Format("02/01/2006"))             // 17/10/2024
	fmt.Println(t.Unix())                           // 1729177445
	fmt.Println(t.UnixMilli())                      // 1729177445123
	fmt.Println(t.Zone())                           // UTC 0

	fmt.Println("Parsing time")
	t3, err := time.Parse("2006-01-02 15:04:05", "2024-10-17 15:04:05")
	fmt.Println(t3, err) // 2024-10-17 15:04:05 +0000 UTC <nil>

	fmt.Println("Working with Time Zones")
	fmt.Println(t.Location()) // UTC

	loc, _ := time.LoadLocation("America/New_York")
	t4 := t.In(loc)
	fmt.Println(t4) // 2024-10-17 11:04:05.123456789 -0400 EDT

	fmt.Println(t.In(time.Local)) // 2024-10-17 11:04:05.123456789 -0400 EDT
	fmt.Println(t.In(time.UTC))   // 2024-10-17 15:04:05.123456789 +0000 UTC
	loc, _ = time.LoadLocation("Asia/Kolkata")
	fmt.Println(t.In(loc)) // 2024-10-17 20:34:05.123456789 +0530 IST

	fmt.Println("Working with Duration")
	t5 := time.Date(2024, time.October, 17, 15, 4, 5, 123456789, time.UTC)
	t6 := time.Date(2024, time.October, 17, 15, 5, 6, 123456789, time.UTC)
	fmt.Println(t6.Sub(t5))            // 1m1s
	fmt.Println(t6.Sub(t5).Seconds())  // 61
	fmt.Println(t6.Sub(t5).Minutes())  // 1
	fmt.Println(time.Since(t6))        // 1m1s
	fmt.Println(t6.Sub(t5).Abs())      // 1m1s
	fmt.Println(t1.Round(time.Minute)) // 2024-10-17 15:04:00 +0000 UTC

}
