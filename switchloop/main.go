package main

import (
	"fmt"
	"math/rand"
)

func main() {

	number := rand.Intn(200) - 100
	switch {
	case number < 0:
		fmt.Println("Negative number")
	case number == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive number")
	}
label: //code comes eher so infinite

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	numbers := []int{2, 2, 3, 4, 5}

	for _, value := range numbers {
		if value == 3 {
			goto label
		}
		fmt.Printf("Index: %d, Value: %d\n", value)
	}
}
