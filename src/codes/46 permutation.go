package codes

func Permute(nums []int) [][]int {
	ans:=[][]int{}
	l:=len(nums)
	if l<1{return ans}
	if l==1{return [][]int{{nums[0]}}}
	t:=0
	for i:=0;i<l;i++{
		tmp:=make([]int,l-1)
		copy(tmp,nums[:i])
		copy(tmp[i:],nums[i+1:])
		ans=append(ans,Permute(tmp)...)
		for ;t<len(ans);t++{
			ans[t]=append(ans[t],nums[i])
		}
	}

	return ans
}
