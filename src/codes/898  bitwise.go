package codes

import "fmt"

func SubarrayBitwiseORs(A []int) int {
	ans:=make(map[int]int)
	cur:=make(map[int]int)
	cur[0]=1
	for _,v:=range A{
		cur2:=make(map[int]int)
		for _,k:= range cur{
			cur2[k|v]=1
			ans[k|v]=1
		}
		cur2[v]= 1
		ans[v]=1
		cur=cur2


		}
	fmt.Println(ans)
	return len(ans)
}
func cc(a []int)int{
	tmp:=0
	for _,v :=range a{
		tmp|=v
	}
	return tmp
}
