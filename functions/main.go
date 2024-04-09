package main

import (
	"fmt"
)

func main() {
	a()
	fmt.Println(adder(3,2))
	ans,_:=proadder(3,2,1,11,1,1,1,1,1,1,1,1,1)
	fmt.Println(ans)

}
func adder(val1 int,val2 int)int {
	return val1+val2
}
func proadder(val1 ...int)(int,string) {
	total:=0
	for _,val := range val1{
		total+=val
	}
	return total,"lmao"
}
func a() {
	fmt.Println("a")
}