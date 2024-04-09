package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Format("2006-02-01 Monday Jan")
	fmt.Println(start)
}
