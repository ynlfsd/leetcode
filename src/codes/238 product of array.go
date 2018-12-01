package codes

func productExceptSelf(nums []int) []int {
	sum:=1
	zero:=0
	for _,k:=range nums{
		if k!=0{sum*=k}else{zero++}
	}
	ans:=make([]int,len(nums))
	if zero>=2{return ans}
	for i:=range ans{
		if zero==1{
			if nums[i]==0{ans[i]=sum}
		}
		if zero==0{
			ans[i]=sum/nums[i]
		}
	}
	return ans
}
