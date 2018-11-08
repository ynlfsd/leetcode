package codes
//
//import (
//	"math"
//	"sort"
//)
//
//func ThreeSumClosest(nums []int, target int) int {
//	ans:=math.MaxInt32
//	sort.Ints(nums)
//	for i:=0;i<len(nums)-2;i++{
//		for j,k:=i+1,len(nums)-1;j<k;{
//			tmp:=target-(nums[i]+nums[j]+nums[k])
//			if abs(ans)>abs(tmp){
//				ans=tmp
//			}
//			if tmp>0{
//				j++
//			}else if tmp<0{
//				k--
//			}else{
//				return target
//
//			}
//
//		}
//	}
//	return target-ans
//}
//func abs(i int)int{
//	if i<0{
//		return i*(-1)
//	}
//	return i
//}