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

	a:=[]int{6,1,3,2,4,7}

	fmt.Println(codes.MaxProfit3(a))

}

func slincetest(ca *[]int){
	*ca=append(*ca,5)

}