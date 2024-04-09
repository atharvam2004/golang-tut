package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content:="hello world"
	file,err:=os.Create("./mygofile.txt")
	// if err!=nil{
	// 	fmt.Print(err)
	// }
	checkNilErr(err)
	length,err:=io.WriteString(file,content)
	checkNilErr(err)

	fmt.Println(length)
	defer file.Close()
	readfile("./mygofile.txt")

}
func readfile(filename string){
	data,err:=os.ReadFile(filename)
	checkNilErr(err)
	fmt.Print(string(data))

}
func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}