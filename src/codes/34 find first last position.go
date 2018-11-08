package codes

func SearchRange(nums []int, target int) []int {
	ans:=[]int{}
	l:=len(nums)
	start,end,mark:=0,l-1,-1
	for start<=end{
		if nums[start+(end-start)/2]==target{
			mark=start+(end-start)/2
			break
		}
		if nums[start+(end-start)/2]<target{
			start=start+(end-start)/2+1
		}else {
			end=start+(end-start)/2-1
		}

	}
	if mark==-1{
		return []int{-1,-1}
	}
	k:=1
	for ;mark-k>=0;k++{
		if nums[mark]!=nums[mark-k]{
			break
		}
	}
	ans=append(ans,mark-k+1)
	for k=1;mark+k<l;k++{
		if nums[mark]!=nums[mark+k]{
			break
		}
	}
	ans=append(ans,mark+k-1)
	return ans
}

