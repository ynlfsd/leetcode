package codes

func numberOfLines(widths []int, S string) []int {
	l:=1
	count:=0
	for _,k:=range S{
		count+=widths[k-'a']
		if count>100{
			l++
			count=widths[k-'a']
		}
	}
	if count==0{return []int{l-1,100}}
	return []int{l,count}
}
