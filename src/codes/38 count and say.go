package codes

import "strconv"

func CountAndSay(n int) string {
	t:="1"
	for i:=1;i<n;i++{
		t=count(t)

	}
	return t
}
func count(s string)string{
	count:=1
	ans:=[]byte{}
	for i:=0;i<len(s);i++{
		if i!=len(s)-1&&s[i+1]==s[i]{
			count++
		}else{
			ans=append(ans,strconv.Itoa(count)...)
			ans=append(ans,s[i])

			count=1
		}
	}

	return string(ans)
}