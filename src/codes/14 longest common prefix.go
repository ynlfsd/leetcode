package codes

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs)==0||len(strs[0])==0{return ""}
	if len(strs)==1 {return strs[0]}
	i:=1
	for prefix(strs[0][:i],strs){
		i++
		if i>=len(strs[0]){break}
	}
	if i==1{return ""}
	return strs[0][:i-1]
}
func prefix(s1 string,s2 []string)bool{
	for _,s:=range s2{
		if !strings.HasPrefix(s,s1){
			return false
		}
	}


	return true
}