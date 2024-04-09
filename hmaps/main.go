package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int) // Creates an empty map
    m["a"] = 1
    m["b"] = 2
	delete(m,"a")
    fmt.Println(m["a"]) // Output: map[a:1 b:2]
	for  k, v := range m {
		fmt.Println(k,v) // Output: map[a:1 b:2]
	}
}
