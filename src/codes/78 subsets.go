package codes

import "fmt"

func Subsets(nums []int) [][]int {
	ans:=[][]int{}
	for i,k:=range nums{
		ans=append(ans,subset([][]int{{k}},nums[i+1:])...)
	}
	ans=append(ans,[]int{})
	return ans
}
func subset(n [][]int,nums []int)[][]int{
	if len(nums)==0 {return n}
	tmp:=[][]int{}
	for i:=0;i<len(n);i++{
		t:=make([]int,len(n[i]))
		copy(t,n[i])
		tmp=append(tmp,t)
	}

	fmt.Println(tmp)
	for i:=0;i<len(tmp);i++{
		tmp[i]=append(tmp[i],nums[0])
	}
	n=append(n,tmp...)
	return subset(n,nums[1:])
}
