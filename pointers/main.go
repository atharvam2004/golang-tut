package main

import (
	"fmt"
)

func main() {
	fmt.Println("pointers")
	var num int = 23
	var ptr *int = &num
	var sum=*ptr+10;
	fmt.Println("Value of ptr", ptr, "nalla", *ptr, "here", sum)
}
