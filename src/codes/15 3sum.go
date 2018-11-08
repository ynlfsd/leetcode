package codes

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	pos:=0
	ans:=make([][]int,0)
	m:=make(map[int][]int)
	for i,v:=range nums{
		if v>=0{
			pos=i
			break
		}
	}
	if pos==0&&len(nums)>3{
		if nums[0]+nums[1]+nums[0]==0{
			return [][]int{{0,0,0}}
		}
		return [][]int{}
	}
	if nums[len(nums)-1]<0{

		return [][]int{}
	}
	if nums[len(nums)-1]==0&&len(nums)>3{
		if nums[len(nums)-3]==0{
			return [][]int{{0,0,0}}
		}
		return [][]int{}
	}
	for i:=0;i<pos+1;i++{
		if nums[i]>0{break}
		for j:=i+1;j<len(nums);j++{
			k:=pos
			if j>pos-1{k=j+1}
			tmp:=nums[i]+nums[j]
			for ;k<len(nums);k++{
				if tmp+nums[k]==0{
					fmt.Println(nums[i])


					if _,ok:=m[nums[i]];!ok||!contains(nums[k],m[nums[i]]){

						m[nums[i]]=append(m[nums[i]],nums[k])
						ans=append(ans,[]int{nums[i],nums[j],nums[k]})

					}
					break
				}
				if tmp+nums[k]>0{break}
			}
		}
	}
	return ans
}
func contains(a int ,b []int)bool{
	for _,k:=range b{
		if k==a{
			return true
		}
	}

	return false
}