package codes


func containsDuplicate(nums []int) bool {
	l:=len(nums)
	if l<2{return false}
	m:=make(map[int]int)
	for _,k:=range nums{
		if _,ok:=m[k];ok{
			return true
		}else{
			m[k]=1
		}
	}
	return false
}