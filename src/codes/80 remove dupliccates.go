package codes

import "fmt"

func RemoveDuplicates80(nums []int) int {
	l:=len(nums)
	if l==0{return 0}
	if l==1{return 1}
	index:=0
	count:=1
	for i:=1;i<l;i++{
		for i<l&&nums[i]==nums[i-1]{
			count++
			if count<=2{
				nums[i-index]=nums[i]
			}
			i++
		}
		if count==1{
			nums[i-index]=nums[i]
		}
		if count>=2{
			index+=count-2
			i--
		}
		count=1
	}
	fmt.Println(nums)
	return l-index
}