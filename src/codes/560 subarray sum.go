package codes

import "fmt"

func SubarraySum(nums []int, k int) int {
	ans:=[]int{}
	count:=0
	sum:=0
	l:=len(nums)
	if l==0{return 0}
	if l==1{
		if nums[0]==k{return 1}else{return 0}
	}
	for _,i:=range nums{
		sum+=i
		if sum==k{count++}
		ans=append(ans,sum)
	}
	fmt.Println(count)
	fmt.Println(ans)
	count+=subSum(ans[1:],nums[1:],k+nums[0])
	return count
}
func subSum(n,nums []int,k int)int{
	l:=len(n)
	if l==0{
		return 0
	}
	count:=0
	for _,i:=range n{
		if i==k{count++}
	}
	count+=subSum(n[1:],nums[1:],k+nums[0])
	return count

}