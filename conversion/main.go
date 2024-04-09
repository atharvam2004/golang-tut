package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input2, _ := reader.ReadString('\n')
	fmt.Println(input2)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	fmt.Println(err)

	numRating = numRating + 1
	fmt.Println(numRating)
}