package codes

func findTheDifference(s string, t string) byte {
	ans:=byte(0)
	for i:=0;i<len(s);i++{
		ans^=s[i]
	}
	for i:=0;i<len(t);i++{
		ans^=t[i]
	}
	return ans
}
