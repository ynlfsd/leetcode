package codes

func singleNumber(nums []int) int {
	l:=len(nums)
	if l==0 {return 0}
	ans:=0
	for _,k:=range nums{
		ans^=k
	}
	return ans

}