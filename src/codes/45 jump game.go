package codes

func Jump(nums []int) int {
	l:=len(nums)
	if l<2{return 0}
	pre,cur,i,res:=0,0,0,0
	for cur<l-1{
		res++
		pre=cur
		for ;i<=pre;i++{
			if cur<i+nums[i]{cur=i+nums[i]}
		}
	}
	return res
}
