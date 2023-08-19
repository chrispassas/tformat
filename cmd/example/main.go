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
