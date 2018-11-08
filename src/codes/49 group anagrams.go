package codes

func GroupAnagrams(strs []string) [][]string {
	m:=make(map[int][]string)
	ans:=[][]string{}
	for _,k:=range strs{
		if _,ok:=m[hash(k)];!ok{
			m[hash(k)]=[]string{k}
		}else{
			m[hash(k)]=append(m[hash(k)],k)
		}
	}
	for _,v:=range m{
		ans=append(ans,v)
	}
	return ans
}
func hash(str string)int{
	multi:=1
	sum:=0
	for i:=0;i<len(str);i++{
		multi*=int((str[i]))
		sum+=int((str[i]-'a'))
	}
	return multi+sum*13
}
