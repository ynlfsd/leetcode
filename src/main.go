package main

import (
	"codes"
	"fmt"
)

type ptest struct {
	a int
	b int
}
func main(){

	a:="123"
	b:="456"
	fmt.Println(codes.Multiply(a,b))

}

func slincetest(ca *[]int){
	*ca=append(*ca,5)

}