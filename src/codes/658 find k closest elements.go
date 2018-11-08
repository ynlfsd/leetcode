package codes

import (
	"sort"
)

func FindClosestElements(arr []int, k int, x int) []int {
	l:=len(arr)
	if k==l{return arr}
	ans:=[]int{}
	switch  {
	case x<=arr[0]:
		count:=0
		for i:=0;i<l;i++{

			ans=append(ans,arr[i])
			count++
			if count==k{
				sort.Ints(ans)
				return ans}
		}
	case x>=arr[l-1]:
		count:=0
		for i:=l-1;i>=0;i--{

			ans=append(ans,arr[i])
			count++
			if count==k{
				sort.Ints(ans)
				return ans}
		}
	default:
		i:=bsearch(arr,x)
		j:=i-1
		count:=0

		for j>=0&&i<=l-1{
			if x-arr[j]<=arr[i]-x{

				ans=append(ans,arr[j])
				count++
				if count==k{
					sort.Ints(ans)
					return ans}
				j--
			}else {

				ans=append(ans,arr[i])
				count++
				if count==k{
					sort.Ints(ans)
					return ans}
				i++
			}
		}

		if j<0&&i<l{
			for i<l{
				ans=append(ans,arr[i])

				i++
				count++
				if count==k{
					sort.Ints(ans)
					return ans}
			}
		}else{
			for j>=0{
				ans=append(ans,arr[j])

				j--
				count++
				if count==k{
					sort.Ints(ans)
					return ans}
			}
		}
	}
	sort.Ints(ans)
	return ans
}
func bsearch(arr []int,t int)int{
	i:=0
	j:=len(arr)-1
	for i<j{
		tmp:=(i+j)/2
		if arr[tmp]==t{
			return tmp
		}else if arr[tmp]>t{
			j=tmp-1
		}else {
			i=tmp+1
		}
	}
	return i
}
