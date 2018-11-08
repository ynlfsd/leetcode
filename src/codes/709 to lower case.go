package codes

func ToLowerCase(str string) string {
	l:=len(str)
	if l==0 {return str}
	ans:=[]byte(str)
	for i:=0;i<l;i++{
		if ans[i]>='A'&&ans[i]<='Z'{
			ans[i]+=32

		}
	}
	return string(ans)
}
