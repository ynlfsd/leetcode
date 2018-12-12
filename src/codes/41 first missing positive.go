package codes

func firstMissingPositive(nums []int) int {


	m:=make(map[int]bool)
	mark:=1
	for _,k:=range nums{
		if k<0{continue}
		m[k]=true
		if k==mark{
			for{
				mark++
				if _,ok:=m[mark];!ok{
					break
				}
			}
		}
	}
	return mark
}