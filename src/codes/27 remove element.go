package codes


func removeElement(nums []int, val int) int {
	l:=len(nums)
	index:=0
	for i:=0;i<l;i++{
		if nums[i]!=val{
			nums[index]=nums[i]
			index++
		}
	}
	return index
}
