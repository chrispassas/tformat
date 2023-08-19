# tformat
This module provides a number of time format strings to make it easier to use Go's time.Format().

## Example
Print out results of different example formats

```go
package main

import (
	"fmt"
	"time"

	"github.com/chrispassas/tformat"
)

func main() {

	currentTime := time.Now()

	fmt.Printf("tformat.YYYY_MM_DD:%s\n", currentTime.Format(tformat.YYYY_MM_DD))
	fmt.Printf("tformat.PGTIMESTAMP:%s\n", currentTime.Format(tformat.PGTIMESTAMP))
	fmt.Printf("tformat.YYYYMMDD:%s\n", currentTime.Format(tformat.YYYYMMDD))
	fmt.Printf("tformat.Timezone:%s\n", currentTime.Format(tformat.Timezone))
	fmt.Printf("tformat.WeekDay:%s\n", currentTime.Format(tformat.WeekDay))
}
```

## Output
```bash
tformat.YYYY_MM_DD:2023-08-19
tformat.PGTIMESTAMP:2023-08-19 14:55:20.753279
tformat.YYYYMMDD:20230819
tformat.Timezone:EDT
tformat.WeekDay:Saturday
```


