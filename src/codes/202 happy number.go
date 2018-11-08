package codes

import "fmt"

func IsHappy(n int) bool {
	if n==1||n==7{return true}
	if n<10{return false}
	ans:=0
	for n>0{
		ans=ans+(n%10)*(n%10)
		n/=10
	}
	fmt.Println(ans)
	if ans==1{return true}
	return IsHappy(ans)
}
