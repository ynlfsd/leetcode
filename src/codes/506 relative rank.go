package codes

import (
	"fmt"
	"sort"
	"strconv"
)

func FindRelativeRanks(nums []int) []string {
	var s =make([]string,len(nums))
	m:=make(map[int]int)
	for v,k:=range nums{
		fmt.Printf("v:%d k:%d\n",v,k)
		m[k]=v
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	for i,v:=range nums{
		switch i{
		case 0:s[m[v]]="Gold Medal"
		case 1:s[m[v]]="Silver Medal"
		case 2:s[m[v]]="Bronze Medal"
		default:s[m[v]]=strconv.Itoa(i+1)

		}
		//s[m[v]]=strconv.Itoa(i+1)
		fmt.Println(s[m[v]])
	}

	return s
}
