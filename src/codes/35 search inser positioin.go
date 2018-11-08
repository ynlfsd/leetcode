package codes

func searchInsert(nums []int, target int) int {
	if target<=nums[0]{return 0}
	if target>nums[len(nums)-1]{return len(nums)}
	for i,k:=range nums{
		if target<=k{
			return i
		}

	}
	return len(nums)
}
