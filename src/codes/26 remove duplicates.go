package codes


func removeDuplicates(nums []int) int {
	l:=len(nums)
	if l<2{return l}
	n:=1
	for i:=1;i<l;i++{
		if nums[i]==nums[i-1]{
			continue

		}else{
			nums[n]=nums[i]
			n++
		}
	}
	return n
}