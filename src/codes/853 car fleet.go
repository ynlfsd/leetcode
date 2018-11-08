package codes

import (
	"sort"
)

func CarFleet(target int, position []int, speed []int) int {
	l:=len(position)
	if l==1{return 1}
	ans:=l
	m:=make(map[int]float64,l)
	for i,k:= range position{
		m[k]=float64(target-k)/float64(speed[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(position)))
	for i:=0;i<l-1;i++{
		for j:=i+1;j<l;j++{
			if m[position[i]]>=m[position[j]]{

				ans--

			}else{
				i=j-1
				break
			}
			if j==l-1{i=j-1}
		}
	}
	return ans
}
func canCatchup(m map[int]float64,i,j,target int)bool{

	if m[i]>=m[j]{
		return true
	}
	return false
}
