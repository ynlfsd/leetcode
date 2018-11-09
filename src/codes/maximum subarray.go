package codes

func MaxSubArray(nums []int) int {
	l:=len(nums)
	if l==1{return nums[0]}
	max:=nums[0]

	for i:=0;i<l;i++{
		if nums[i]<=0{
			if max<nums[i]{
				max=nums[i]
			}
			continue}
		sum:=0
		j:=i
		for ;j<l;j++{
			if sum+nums[j]<0{break}
			sum+=nums[j]
			if max<sum{max=sum}
		}
		i=j
	}
	return max
}
