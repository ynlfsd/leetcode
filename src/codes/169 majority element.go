package codes

func MajorityElement(nums []int) int {
	l:=len(nums)
	if l==1{return nums[0]}
	i,ans:=1,nums[0]
	for _,k:=range nums[1:]{
		if k==ans{
			i++
		}else{
			i--
			if i==0{
				ans=k
			}
		}
	}
	return ans
}
