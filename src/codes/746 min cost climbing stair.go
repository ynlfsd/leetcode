package codes

import "fmt"

func MinCostClimbingStairs(cost []int) int {
	l:=len(cost)
	ret:=make([]int,l)
	ret[0]=cost[0]
	ret[1]=cost[1]
	for i:=2;i<l;i++{
		ret[i]=cost[i]+min(ret[i-1],ret[i-2])
	}
	fmt.Println(ret)
	return min(ret[l-1],ret[l-2])
}
func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}
