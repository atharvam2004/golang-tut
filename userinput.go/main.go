package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// input,_:=reader.ReadString('\n')
	// fmt.Println("haha",input)
	input2,_:=reader.ReadString('\n')//_ can be declared and not used, other car names like err should be used
	fmt.Println(input2)
}
