package codes


func containsNearbyDuplicate(nums []int, k int) bool {
	m:=make(map[int]int)
	for v,key:=range nums{
		if a,ok:=m[key];ok{
			if v-a<=k{
				return true
			}

		}
		m[key]=v
	}


	for i:=0;i<len(nums);i++{
		for j:=i+1;j<i+k+1&&j<len(nums);j++{
			if nums[i]==nums[j]{
				return true
			}
		}
	}





	return false
}