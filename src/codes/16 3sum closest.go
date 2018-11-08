package codes

import (
	"sort"
	"fmt"
)

func ThreeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	sum:=0
	for i:=0;i<=2;i++{
		sum+=nums[i]
	}
	ans:=check(nums[:3],nums[3:],sum,target)

	return ans
}
func check(a,b []int , sum ,target int)int{
	s:=0
	e:=len(b)-1
	for{
		switch  {
		case sum-target<0:
			tmp:=[]int{}
			for j:=0;j<3;j++{
				tmp=append(tmp,(sum-a[j]+b[e]-target)*(sum-a[j]+b[e]-target))
			}
			t:=findmin(tmp,(sum-target)*(sum-target))
			if t!=-1{
				sum=sum-a[t]+b[e]
				a[t]=b[e]
				sort.Ints(a)
			}
			e-=1
		case sum-target>0:
			tmp:=[]int{}
			for j:=0;j<3;j++{
				fmt.Println(s)
				tmp=append(tmp,(sum-a[j]+b[s]-target)*(sum-a[j]+b[s]-target))
			}
			t:=findmin(tmp,(sum-target)*(sum-target))
			if t!=-1{
				sum=sum-a[t]+b[s]
				a[t]=b[e]
				sort.Ints(a)
			}
			s+=1
		case sum-target==0:return sum
		}
		if s>=e{
			return sum
		}
	}



	return sum
}

func findmin(a []int, b int)int{
	ans:=-1
	for i,k:=range a{
		if k<b{
			ans=i
			b=k
		}
	}
	return ans
}