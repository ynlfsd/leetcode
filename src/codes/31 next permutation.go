package codes
//
//
//func NextPermutation(nums []int)  {
//	l:=len(nums)
//	if l<2{return}
//	var top,low int
//	mark:=false
//	for i:=l-1;i>0;i--{
//		if nums[i]>nums[i-1]{
//			mark=true
//			break
//		}
//
//	}
//	if i==1&&!mark{
//		for i:=0;i<l/2;i++{
//			nums[i],nums[l-i-1]=nums[l-i-1],nums[i]
//
//		}
//		return
//	}
//	if i==l-1{
//
//	}
//	if i>=1{
//		changed(nums[i-1:])
//	}
//
//
//
//
//
//}
//func changed(nums []int){
//	tmp:=nums[l-1]
//	for i:=l-1;i>0;i--{
//		nums[i]=nums[i-1]
//	}
//	nums[0]=tmp
//}
