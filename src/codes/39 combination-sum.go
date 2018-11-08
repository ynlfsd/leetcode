package codes

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	l:=len(candidates)
	ans:=[][]int{}
	if l==0{return [][]int{}}
	sort.Ints(candidates)
	if candidates[0]>target{return [][]int{}}
	if candidates[0]==target{return [][]int{{target}}}
	for i:=0;i<l;i++{
		if candidates[i]<target{
			ans=append(ans,combina(candidates[i:],candidates[i],target)...)
		}else if candidates[i]==target{
			ans=append(ans,[]int{candidates[i]})
		}

	}
	return ans
}
func combina(ca []int,sum int,target int)[][]int{
	ans:=[][]int{}
	for i:=0;i<len(ca);i++{
		if sum+ca[i]<target{
			ans=append(ans,combina(ca[i:],sum+ca[i],target)...)
		}else if sum+ca[i]==target{
			ans=append(ans,[]int{ca[i]})
		}
	}
	if len(ans)>0{
		for i:=0;i<len(ans);i++{
			ans[i]=append(ans[i],ca[0])
		}
	}

	return ans

}