package codes

import "fmt"

func MMajorityElement(nums []int) []int {
	l:=len(nums)
	if l<3 {return []int{nums[0]}}
	ans:=[]int{}
	num1,num2,count1,count2:=nums[0],nums[0],1,1
	for _,k:=range nums[1:]{
		if k==num1{
			count1++
		}else if k==num2{
			count2++
		}else if count1==0{
			num1=k
			count1=1
		}else if count2==0{
			num2=k
			count2=1
		}else{
			count1--
			count2--
		}
	}
	fmt.Println(num1)
	fmt.Println(num2)
	count1,count2=0,0
	for _,k:=range nums{
		if k==num1{
			count1++
		}else if k==num2{
			count2++
		}
	}
	if count1>l/3{ans=append(ans,num1)}
	if num1==num2{return ans}
	if count2>l/3{ans=append(ans,num2)}
	return ans
}
