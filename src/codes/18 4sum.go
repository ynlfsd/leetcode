package codes

import "sort"

func FourSum(nums []int, target int) [][]int {
	if len(nums)==4{
		if nums[0]+nums[1]+nums[2]+nums[3]==target{
			return [][]int{{nums[0],nums[1],nums[2],nums[3]}}
		}
		return [][]int{}

}
	sort.Ints(nums)
	ans:=[][]int{}
	ansi:=-1
	ansj:=0
	ansk:=0
	ansl:=0
	for i:=0;i<len(nums)-3;i++{
		for j:=i+1;j<len(nums)-2;j++{
			if i!=0&&nums[i]==nums[i-1]&&nums[j]==nums[j-1]{
				continue
			}
			if j!=i+1&&nums[j]==nums[i+1]{
				continue
			}
			if j!=i+1&&nums[j]==nums[j-1]{
				continue
			}
			if i!=0&&nums[i]==nums[i-1]{
				continue
			}
			if i!=0&&nums[i]==nums[ansi]&&nums[j]==nums[ansj]{
				continue
			}
			for k,l:=j+1,len(nums)-1;k<l;{
				sum:=nums[i]+nums[j]+nums[k]+nums[l]
				if sum==target{
					if ansi==-1{
						ans=append(ans,[]int{nums[i],nums[j],nums[k],nums[l]})
						ansi=i
						ansj=j
						ansk=k
						ansl=l
					}else if !(nums[i]==nums[ansi]&&nums[j]==nums[ansj]&&nums[k]==nums[ansk]&&nums[l]==nums[ansl]){
						ans=append(ans,[]int{nums[i],nums[j],nums[k],nums[l]})
						ansi=i
						ansj=j
						ansk=k
						ansl=l

					}

					k++
					l--
				}else if sum>target{
					l--
				}else{
					k++
				}
			}
		}
	}
	return ans
}
