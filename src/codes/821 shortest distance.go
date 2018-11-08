package codes

import (
	"fmt"
	"strings"
)

func ShortestToChar(S string, C byte) []int {
	l:=len(S)
	ans:=make([]int,l)
	for i:=0;i<l;i++{
		ans[i]=l
	}
	start:=0
	for {
		tmp:=strings.IndexByte(S[start:],C)
		fmt.Println(tmp)
		if tmp==-1{break}
		distance(ans,start+tmp)
		start=start+tmp+1
	}
	return ans
}
func distance(ans []int, i int){
	left,right:=i-1,i+1
	ans[i]=0
	for left>=0{
		if ans[left]>ans[left+1]+1{
			ans[left]=ans[left+1]+1
			left=left-1
		}else{break}
	}
	for right<len(ans){
		if ans[right]>ans[right-1]+1{
			ans[right]=ans[right-1]+1
			right=right+1
		}else{break}
	}
}
