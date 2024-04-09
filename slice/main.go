package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to array in golang")
	var fruitlist [3]string
	fruitlist[0] = "69"
	fruitlist[1] = "69"
	fmt.Println(fruitlist)
	// Declare an array of integers with values
	numbers := []int{1, 2, 3, 4, 5}
	numbers=append(numbers, 1,2)
	// Print the array
	fmt.Printf("Array: %T", numbers)
	num2:=append(numbers[1:3],24)
	sort.Ints(num2)
	fmt.Println("Array:",num2)
	fmt.Println(sort.IntsAreSorted(num2))
	num3:=append(numbers[:1],numbers[2:]...)
	fmt.Println(num3)
}
