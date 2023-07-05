---
title: Markdown to Slides
author: Piyush Patel
---

# Here is the first slide

- The first bullet point.
- The second one.
    - And the third.
    - Another slide 

# This will start a new slide.

- New bullet point.
- etc.
- etc.

## Code

```go
package main

import "fmt"

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)
	var numbers2 [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(numbers2)
}
```

```shell
$ go run main.go
[1 2 3 4]
[1 2 3 4]
```

