package codes

import "strings"

func lengthOfLastWord(s string) int {
	if len(s)==0{return 0}
	ans:=strings.Fields(s)
	if len(ans)<=1{return 0}
	return len(ans[len(ans)-1])
}
