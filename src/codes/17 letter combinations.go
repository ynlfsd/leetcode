package codes

func LetterCombinations(digits string) []string {
	m:=make(map[uint8][]byte)
	m['2']=[]byte{'a','b','c'}
	m['3']=[]byte{'d','e','f'}
	m['4']=[]byte{'g','h','i'}
	m['5']=[]byte{'j','k','l'}
	m['6']=[]byte{'m','n','o'}
	m['7']=[]byte{'p','q','r','s'}
	m['8']=[]byte{'t','u','v'}
	m['9']=[]byte{'w','x','y','z'}
	ans:=[][]byte{}
	for i:=0;i<len(digits);i++{

		if digits[i]=='1'{continue}
		if i==0{
			for _,v :=range m[digits[i]]{
				ans=append(ans,[]byte{v})

			}

		}else{

			l:=len(ans)
			for t:=0;t<len(m[digits[i]])-1;t++{
				for d:=0;d<l;d++{
					tmp:=make([]byte,len(ans[d]))

					copy(tmp,ans[d])

					ans=append(ans,tmp)
				}


			}


			for t:=0;t<len(ans);t++{
					c:=m[digits[i]][t/l]
					ans[t]=append(ans[t],c)


			}

		}

	}

	s:=[]string{}
	for _,t:=range ans{

		s=append(s,string(t))
	}
	return s
}
