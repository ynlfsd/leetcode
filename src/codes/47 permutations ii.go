package codes

import "sort"

func PermuteUnique(nums []int) [][]int {
	ans:=[][]int{}
	l:=len(nums)
	if l<1{return ans}
	if l==1{return [][]int{{nums[0]}}}
	sort.Ints(nums)
	t:=0
	for i:=0;i<l;i++{
		if i!=0&&nums[i-1]==nums[i]{
			continue
		}
		tmp:=make([]int,l-1)
		copy(tmp,nums[:i])
		copy(tmp[i:],nums[i+1:])
		ans=append(ans,PermuteUnique(tmp)...)
		for ;t<len(ans);t++{
			ans[t]=append(ans[t],nums[i])
		}
	}

	return ans
}