package codes

func uniqueMorseRepresentations(words []string) int {
	m:=map[rune]string{'a':".-",'b':"-...",'c':"-.-.",'d':"-..",'e':".",'f':"..-.",'g':"--.",'h':"....",'i':"..",'j':".---",'k':"-.-",'l':".-..",'m':"--",'n':"-.",'o':"---",'p':".--.",'q':"--.-",'r':".-.",'s':"...",'t':"-",'u':"..-",'v':"...-",'w':".--",'x':"-..-",'y':"-.--",'z':"--.."}
	ans:=make(map[string]int)
	ret:=0
	for _,s:=range words{
		tmp:=[]byte{}
		for _,i:=range s{
			tmp=append(tmp,m[i]...)
		}
		if _,ok:=ans[string(tmp)];!ok{
			ans[string(tmp)]=1
			ret++
		}
		tmp=nil
	}
	return ret
	}
