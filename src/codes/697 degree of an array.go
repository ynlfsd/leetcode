package codes

import "fmt"

func FindShortestSubArray(nums []int) int {
	num:=make([]int,50000)
	for _,i:=range nums{
		num[i]+=1
	}
	max:=0
	res:=[]int{}
	for _,i:=range num{
		if max<i{max=i}
	}
	fmt.Println(max)
	for k,i:=range num{
		if i==max{
			res=append(res,k)
		}
	}
	fmt.Println(res)
	l:=len(nums)
	for _,i:=range res{
		j:=0
		k:=len(nums)-1
		for{
			fmt.Println(j,k)
			if nums[j]==i&&nums[k]==i{break}
			if nums[j]!=i{j++}
			if nums[k]!=i{k--}
		}
		if l>k-j+1{l=k-j+1}
	}
	return l
}
