package codes

func search(nums []int, target int) int {
	l:=len(nums)
	start,end:=0,l-1
	for start<=end{
		if nums[start+(end-start)/2]==target{
			return start+(end-start)/2
		}
		if nums[start+(end-start)/2]<target{
			if nums[start+(end-start)/2]>nums[start]{
					start=start+(end-start)/2+1
			}else{
				if target>=nums[start]{
					end=start+(end-start)/2-1
				}else{
					start=start+(end-start)/2+1
				}
			}

		}else{
			if nums[start+(end-start)/2]>nums[end]{
				end=start+(end-start)/2-1
			}else{
				if target<=nums[end]{
					start=start+(end-start)/2+1
				}else {
					end=start+(end-start)/2-1
				}

			}
		}
	}
	return -1
}
