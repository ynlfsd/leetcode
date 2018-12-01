package codes

func canJump(nums []int) bool {
	l:=len(nums)
	if l==0{return true}
	for i,k:=range nums{
		if i==l-1{break}
		if k!=0{continue}
		can:=false
		for j:=i-1;j>=0;j--{
			if nums[j]+j>i{
				can=true
				break
			}
		}
		if !can{return false}
	}
	return true
}
