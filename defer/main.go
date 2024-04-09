package main

import "fmt"

func printNumbers() {
	defer fmt.Println("Function end") // This will be executed just before the function returns
	fmt.Println("Printing numbers:")
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i) // This will be deferred and executed in LIFO order
	}
}

func main() {
	defer fmt.Println("Function start")
	printNumbers()
	fmt.Println("Main function end")
}
